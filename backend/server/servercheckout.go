package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hackathon/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var mutex sync.Mutex

func (a *Apiserver) Checkout(res http.ResponseWriter, req *http.Request) {
	//get current customer logged in
	customerID, exist := a.CheckSessionExistCustomer(req)
	if !exist {
		fmt.Println("customer session doesnt exist...")
		return
	} else {
		fmt.Println("customer session exist...", customerID)
	}
	//customerID := "C001" //hardcoded first
	if req.Method == http.MethodPost {
		checkoutDetails := make(map[string]interface{})
		if req.Header.Get("Content-Type") == "application/json" { //Get data passed from client in form format
			res.Header().Set("Content-Type", "application/json")
			err := json.NewDecoder(req.Body).Decode(&checkoutDetails)
			if err != nil {
				fmt.Println("Error with json decoding of req body in login", err)
				res.WriteHeader(http.StatusInternalServerError)
				res.Write([]byte(`{"orderInserted": false, "OrderID": ""}`))
				return
			}

			//ORDER AND ORDER DETAIL
			newOrderID := uuid.NewString()
			orderTime := time.Now()
			orderTimeString := orderTime.Format("2006-01-02 15:04:05")
			totalPrice := checkoutDetails["Total"].(float64)
			deliveryAddress := checkoutDetails["DeliveryAddress"].(string) //hardcoded RiderID!!!
			order := models.Orders{OrderID: newOrderID, CustomerID: customerID, RiderID: "R001", OrderStatus: "CART", OrderEnd: orderTimeString, Total: totalPrice, DeliveryAddress: deliveryAddress}

			//"Meal":[{"ID":"M001","Qty":2, "Price":2.50}]
			var meals []models.OrderDetail
			fetchMeals := checkoutDetails["Meal"].([]interface{})

			for _, meal := range fetchMeals {
				getMealID := meal.(map[string]interface{})["ID"].(string)
				getMealQty := int(meal.(map[string]interface{})["Qty"].(float64)) //This is a common issue when dealing with JSON data in Go, since JSON numbers (even integers) are typically decoded as float64 in Go.
				getMealPrice := meal.(map[string]interface{})["Price"].(float64)
				individualMeal := models.OrderDetail{OrderID: newOrderID, MealID: getMealID, MealQty: getMealQty, MealPrice: getMealPrice}
				meals = append(meals, individualMeal)
			}
			//NEED TO update the discounts qty if order successfully pass thru.
			//also need to check beforehand that qty fits whatever is left

			//Uncomment this code below to test out multiple users :)
			//curl -X POST http://localhost:5001/checkout -H "Content-Type: application/json" -d '{"CustomerID":"C001","Total":5.00, "DeliveryAddress":"Jurong West", "Meal":[{"ID":"M001","Qty":2, "Price":2.50},{"ID":"M002","Qty":2, "Price":2.50}]}'
			//// Simulate multiple customers attempting to purchase items
			//var wg sync.WaitGroup
			//for i := 1; i <= 3; i++ {
			//	wg.Add(1)
			//	go func(id int) {
			//		defer wg.Done()
			//		//purchaseItem(id, 2) // Each customer tries to purchase 2 items
			//		ok := a.ValidateWantedQuantityAgainstStockAndUpdate(meals, res, req)
			//		if !ok {
			//			fmt.Println("Error with validating wanted quantity..dont insert")
			//			res.WriteHeader(http.StatusInternalServerError)
			//			res.Write([]byte(`{"orderInserted": false, "OrderID": ""}`))
			//			return
			//		}
			//	}(i)
			//}
			//
			//wg.Wait() // Wait for all customers to finish their purchases

			ok := a.ValidateWantedQuantityAgainstStockAndUpdate(meals, res, req)
			if !ok {
				fmt.Println("Error with validating wanted quantity..dont insert")
				res.WriteHeader(http.StatusInternalServerError)
				res.Write([]byte(`{"orderInserted": false, "OrderID": ""}`))
				return
			}

			//insert into order first, then orderdetail
			dberr := a.DB.InsertOrder(order)
			if dberr != nil {
				fmt.Println("Error with inserting order in db", dberr)
				res.WriteHeader(http.StatusInternalServerError)
				res.Write([]byte(`{"orderInserted": false, "OrderID": ""}`))
				return
			}

			//now insert into orderdetail
			a.DB.InsertOrderDetail(meals)

			//return headers
			res.WriteHeader(http.StatusCreated)
			res.Write([]byte(`{"orderInserted": true, "OrderID": "` + newOrderID + `"}`))
		}
	}

	if req.Method == http.MethodPut {
		var meals []models.OrderDetail
		fmt.Println("PUT im in here heh")
		// Read the body of the request
		body, err := ioutil.ReadAll(req.Body) //read bytes
		if err != nil {
			log.Println("Error reading request body:", err)
			return
		}

		unmarshalJsonErr := json.Unmarshal(body, &meals)
		if unmarshalJsonErr != nil {
			fmt.Println("error unmarshalling json:", unmarshalJsonErr)
			return
		}
		//fmt.Println("YASS", meals)
		reduceerr := a.DB.ReduceQty(meals)
		if reduceerr != nil {
			fmt.Println("Error with reduce:", reduceerr)
			return
		}
	}
}

// if customer stock is lesser or fits stock, returns true. else customer cannot order
// this should be mutex because only 1 can read at a time ..
// whoever reaches this function first wins!
// crucial to lock it while performing the validation of stock qty too coz if you only lock it when updating the stock and leave the validation outside the critical section, 2 users could both see the stock as available, thus leading to double booking
func (a *Apiserver) ValidateWantedQuantityAgainstStockAndUpdate(meals []models.OrderDetail, res http.ResponseWriter, req *http.Request) bool {
	mutex.Lock()
	defer mutex.Unlock()

	//check wanted qty against Discount table
	ok := a.DB.ValidateDiscountStockQty(meals)
	if !ok {
		return false //cannot buy...exceeds :/
	} else { //can buy and can proceed to update the discount table to minus off qty
		//perform PUT/PATCH METHOD TO UPDATE DISCOUNT TABLE QTY!
		mealData, err := json.Marshal(meals) //we marshal meals data and put inside req.body of put
		if err != nil {
			fmt.Println(err)
			return false
		}

		targetURL := "http://localhost:5001/checkout"

		putReq, err := http.NewRequest(http.MethodPut, targetURL, bytes.NewBuffer(mealData))
		if err != nil {
			fmt.Println("Error creating PUT request", err)
			http.Error(res, "Error creating PUT request", http.StatusInternalServerError)
			return false
		}

		// Copy headers from req to putReq
		putReq.Header = req.Header.Clone()
		putReq.Header.Set("Content-Type", "application/json")

		// Forward PUT request.
		client := &http.Client{}
		putResp, err := client.Do(putReq) //after PUT is done, control is sent back to this post with response from put
		if err != nil {
			fmt.Println("Error forwarding PUT request:", err)
			http.Error(res, "Error forwarding PUT request", http.StatusInternalServerError)
			return false
		}
		defer putResp.Body.Close()
		return true
	}
}
