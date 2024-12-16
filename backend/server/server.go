package server

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"hackathon/database"
	"net/http"
	"strings"
)

type Apiserver struct {
	DB *database.DBClient
}

func NewApiserver(db *database.DBClient) *Apiserver {
	return &Apiserver{
		DB: db,
	}
}

func (a *Apiserver) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", a.Home) // handler "home"
	router.HandleFunc("/vendor/discount", a.VendorDiscount)
	router.HandleFunc("/customer/discount", a.GetCustomerDiscount)                 //show vendor
	router.HandleFunc("/customer/discount/{vid}", a.GetCustomerDiscountIndividual) //show each vendor meals
	router.HandleFunc("/checkout", a.Checkout)                                     //checkout function
	router.HandleFunc("/login", a.Login)
}

func (a *Apiserver) Login(res http.ResponseWriter, req *http.Request) {
	if http.MethodPost == req.Method {
		if req.Header.Get("Content-Type") == "application/json" {
			// Create a map to hold the decoded data
			loginDetails := make(map[string]string)

			// Decode the JSON from the request body into the map
			err := json.NewDecoder(req.Body).Decode(&loginDetails)
			if err != nil {
				fmt.Println("Error with JSON decoding of req body in login:", err)
				return
			}

			fmt.Println("Username:", loginDetails["Username"])
			fmt.Println("Password:", loginDetails["Password"])
			fmt.Println("Vendor:", loginDetails["Role"])
		}

	}
}

func (a *Apiserver) VendorDiscount(w http.ResponseWriter, r *http.Request) {

}

func (a *Apiserver) GetCustomerDiscount(w http.ResponseWriter, r *http.Request) {

}
func (a *Apiserver) GetCustomerDiscountIndividual(res http.ResponseWriter, req *http.Request) {
	//check discount start and discount end. if discount has ended, then we delete that meal from the database.
	if req.Method == http.MethodGet { //GET ALL Meals from Vendors
		//fetch the vendor id
		getVendorID := req.URL.Path
		parts := strings.Split(getVendorID, "/")
		var vendorID string

		if len(parts) > 1 {
			vendorID = parts[len(parts)-1]
		} else {
			return
		}

		//get the all meals from that vendor
		meals, err := a.DB.GetMealFromVendor(vendorID)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "No meals available in vendor"}`))
			return
		}

		//get only meals that are discounted
		discountedMeals, dberr := a.DB.GetDiscountedMealsFromVendor(meals)
		if dberr != nil {
			fmt.Println(dberr)
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "No discounted meals found"}`))
			return
		}

		jsonerr := json.NewEncoder(res).Encode(discountedMeals) //json data returned
		if jsonerr != nil {
			fmt.Println(jsonerr)
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "Error with json conversion. Discounted meals not found"}`))
			return
		}
	}
}

func (a *Apiserver) Checkout(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		checkoutDetails := make(map[string]interface{})
		if req.Header.Get("Content-Type") == "application/json" { //Get data passed from client in form format
			err := json.NewDecoder(req.Body).Decode(&checkoutDetails)
			if err != nil {
				fmt.Println("Error with json decoding of req body in login")
			}

			//fmt.Println(checkoutDetails["data"])

			//ordersMake := models.Orders{OrderID: uuid.NewString(), }

			//data := checkoutDetails["data"].(map[string]interface{})
			//order := data["order"].(map[string]interface{})
			//items := order["items"].([]interface{}) // We expect items to be an array
			//
			//// Loop through the items to access "itemID"
			//for _, item := range items {
			//	itemMap := item.(map[string]interface{})  // type assertion to map for each item
			//	itemID := itemMap["itemID"].(string)      // Access itemID as a string
			//	quantity := itemMap["quantity"].(float64) // Access quantity as a float64 (JSON numbers are decoded as float64)
			//
			//	// Print the item details
			//	fmt.Println("Item ID:", itemID)
			//	fmt.Println("Quantity:", quantity)
			//}
		}

	}
}

//func (a *Apiserver) GetAllUsers(w http.ResponseWriter, r *http.Request) {
//	// iterate all people
//
//	//json.NewEncoder(w).Encode(users)
//}

func (a *Apiserver) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}
