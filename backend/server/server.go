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

		jsonerr := json.NewEncoder(res).Encode(discountedMeals)
		if jsonerr != nil {
			fmt.Println(jsonerr)
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "Error with json conversion. Discounted meals not found"}`))
			return
		}
		res.WriteHeader(http.StatusOK)
	}
}

func (a *Apiserver) Checkout(w http.ResponseWriter, r *http.Request) {

}

//func (a *Apiserver) GetAllUsers(w http.ResponseWriter, r *http.Request) {
//	// iterate all people
//
//	//json.NewEncoder(w).Encode(users)
//}

func (a *Apiserver) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}
