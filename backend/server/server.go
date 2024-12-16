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
	//GET Method -- let the vendor see their “home” page (to edit which meal to go live for discount + quantity, etc)
	if r.Method == http.MethodGet {
		defer r.Body.Close()

		cookie, err := r.Cookie("sessionID")
		if err != nil {
			http.Error(w, "Session cookie not found", http.StatusUnauthorized)
			return
		}
		sessionID := cookie.Value

		var vendorID string
		err = db.QueryRow("SELECT VendorID FROM VendorSessions WHERE SessionID = '%s' AND SessionExpiry > NOW()", sessionID).Scan(&vendorID)
		if err != nil {
			http.Error(w, "Invalid or expired session", http.StatusUnauthorized)
			return
		}

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

	////POST Method -- INSERT the data back to DB once vendor has set which meal to be live for discount and the quantity
	//if r.Method == http.MethodPost {
	//	defer r.Body.Close()
	//
	//	// Read the entire request body
	//	statusData, err := io.ReadAll(r.Body)
	//	if err != nil {
	//		http.Error(w, "Failed to read request body", http.StatusBadRequest)
	//		return
	//	}
	//
	//	fmt.Println("Received data:", string(statusData))
	//
	//	// Unmarshal the request body into a map to retrieve the "discountStatus" field
	//	var requestBody map[string]string
	//	if err := json.Unmarshal(statusData, &requestBody); err != nil {
	//		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
	//		return
	//	}
	//
	//	// Retrieve the "discountStatus" field ("Launch" or "End") from the request body
	//	status, exists := requestBody["discountStatus"]
	//	if !exists || (status != "Launch" && status != "End") {
	//		http.Error(w, "Invalid status, must be 'Launch' or 'End'", http.StatusBadRequest)
	//		return
	//	}
	//
	//	// Call the DiscountStatus method on the DB client
	//	if err := a.DB.DiscountStatus(status); err != nil {
	//		// If DiscountStatus fails, return a 500 error
	//		http.Error(w, "Failed to set discount status: "+err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	// If everything goes well, return a 200 OK with a success message
	//	w.WriteHeader(http.StatusOK)
	//	fmt.Fprintf(w, "Discount status updated to: %v\n", status)
	//}
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
