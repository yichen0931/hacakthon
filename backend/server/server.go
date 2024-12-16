package server

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"hackathon/database"
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
	if req.Method == http.MethodGet { //GET ALL Meals from Vendors
		//getVendorID := req.URL.Path
		//
		//parts := strings.Split(getVendorID, "/")
		//var vendorID string

	}
	//getVendor := r.URL.Path
	//urlPath := req.URL.Path //// Get the full URL path. attendance/<id>
	//
	//parts := strings.Split(urlPath, "/") // Split the URL path by "/"
	//var attendanceID string
	//
	//if len(parts) > 1 { // Assuming the ID is the last part of the URL path
	//	attendanceID = parts[len(parts)-1]
	//} else {
	//	logrus.Info("Attendance ID not provided")
	//	http.Error(res, "Attendance ID not provided", http.StatusBadRequest)
	//	return
	//}
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
