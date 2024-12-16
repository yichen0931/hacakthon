package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hackathon/models"
	"net/http"
	"time"
)

func (a *Apiserver) Checkout(res http.ResponseWriter, req *http.Request) {
	//get current customer logged in
	//customerID, exist := a.CheckSessionExistCustomer(req)
	//if !exist {
	//	fmt.Println("customer session doesnt exist...")
	//	return
	//} else {
	//	fmt.Println("customer session exist...", customerID)
	//}
	customerID := "C002" //hardcoded first
	if req.Method == http.MethodPost {
		checkoutDetails := make(map[string]interface{})
		if req.Header.Get("Content-Type") == "application/json" { //Get data passed from client in form format
			err := json.NewDecoder(req.Body).Decode(&checkoutDetails)
			if err != nil {
				fmt.Println("Error with json decoding of req body in login", err)
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

			//insert into order first, then orderdetail
			dberr := a.DB.InsertOrder(order)
			if dberr != nil {
				fmt.Println("Error with inserting order in db", dberr)
				return
			}

			//now insert into orderdetail
			a.DB.InsertOrderDetail(meals)

			//return headers
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusCreated)
			res.Write([]byte(`{"orderInserted": true, "OrderID": "` + newOrderID + `"}`))
		}
	}
}
