package server

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"hackathon/database"
	"io/ioutil"
	"net/http"
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
	router.HandleFunc("/panda/v1/", a.Home) // handler "home"
	router.HandleFunc("/vendor/discount", a.VendorDiscount)
	router.HandleFunc("/customer/discount", a.GetCustomerDiscount)                 //show vendor
	router.HandleFunc("/customer/discount/{vid}", a.GetCustomerDiscountIndividual) //show each vendor meals
	router.HandleFunc("/checkout", a.Checkout)                                     //checkout function
}

func (a *Apiserver) VendorDiscount(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodGet {
	//	a.DB.VendorGetMeals()
	//}

	//if r.Method == http.MethodPost {
	//	body, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		w.WriteHeader(http.StatusBadRequest)
	//	}
	//	json.Unmarshal(body, &a.DB)
	//	fmt.Println()
	//	a.DB.SetDiscountTime(startTime, endTime)
	//}

	if r.Method == http.MethodPost {
		defer r.Body.Close()

		// Read the entire request body
		statusData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		fmt.Println("Received data:", string(statusData))

		// Unmarshal the request body into a map to retrieve the "discountStatus" field
		var requestBody map[string]string
		if err := json.Unmarshal(statusData, &requestBody); err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
			return
		}

		// Retrieve the "discountStatus" field ("Launch" or "End") from the request body
		status, exists := requestBody["discountStatus"]
		if !exists || (status != "Launch" && status != "End") {
			http.Error(w, "Invalid status, must be 'Launch' or 'End'", http.StatusBadRequest)
			return
		}

		// Call the DiscountStatus method on the DB client
		if err := a.DB.DiscountStatus(status); err != nil {
			// If DiscountStatus fails, return a 500 error
			http.Error(w, "Failed to set discount status: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// If everything goes well, return a 200 OK with a success message
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Discount status updated to: %v\n", status)
	}
}

func (a *Apiserver) GetCustomerDiscount(w http.ResponseWriter, r *http.Request) {

}
func (a *Apiserver) GetCustomerDiscountIndividual(w http.ResponseWriter, r *http.Request) {

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
