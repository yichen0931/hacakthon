package server

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"hackathon/database"
	"hackathon/models"
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

func (a *Apiserver) VendorDiscount(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		http.Error(w, "Session cookie not found", http.StatusUnauthorized)
		return
	}
	sessionID := cookie.Value

	var vendorID string
	err = a.DB.VendorCheckSession(sessionID)
	if err != nil {
		http.Error(w, "Invalid or expired session", http.StatusUnauthorized)
		return
	}

	//GET Method -- let the vendor see their “home” page (to edit which meal to go live for discount + quantity, etc)
	if r.Method == http.MethodGet {
		defer r.Body.Close()

		results, err := a.DB.VendorViewAllMeal(vendorID)
		if err != nil {
			http.Error(w, "Failed to fetch vendor discount details", http.StatusUnauthorized)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}

	//POST Method -- INSERT the data back to DB once vendor has set which meal to be live for discount and the quantity
	if r.Method == http.MethodPost {
		defer r.Body.Close()

		var updatedDiscounts []models.VendorSetDiscount

		if r.Header.Get("Content-Type") == "application/json" { //Get data passed from client in form format
			err := json.NewDecoder(r.Body).Decode(&updatedDiscounts)
			if err != nil {
				fmt.Println("Error with json decoding of req body in login", err)
				http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
				return
			}
			for _, discount := range updatedDiscounts {
				// If MealID is empty, or DiscountedPrice or Quantity is invalid (<= 0),
				// set them as "no discount"
				if discount.MealID == "" || discount.DiscountedPrice < 0 || discount.Quantity < 0 {
					http.Error(w, "Invalid or missing necessary fields", http.StatusBadRequest)
					return
				}

				// If DiscountedPrice or Quantity is 0, consider it as no discount for that meal
				if discount.DiscountedPrice == 0 || discount.Quantity == 0 {
					discount.DiscountedPrice = 0
					discount.Quantity = 0
				}

				//proceed to set the discount to the database.
				err = a.DB.VendorSetDiscount(&discount)
				if err != nil {
					http.Error(w, "Failed to insert discount details", http.StatusInternalServerError)
					return
				}
			}

			// Respond back to the client with success
			w.Header().Set("Content-Type", "application/json")
			response := map[string]string{"message": "Discount updated successfully"}

			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
		}
	}
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

//func (a *Apiserver) GetAllUsers(w http.ResponseWriter, r *http.Request) {
//	// iterate all people
//
//	//json.NewEncoder(w).Encode(users)
//}

func (a *Apiserver) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}
