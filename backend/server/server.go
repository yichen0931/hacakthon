package server

import (
	"encoding/json"
	"fmt"
	"hackathon/database"
	"hackathon/models"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
	cookie, err := r.Cookie("vendorSessionCookie")
	if err != nil {
		http.Error(w, "Session cookie not found", http.StatusUnauthorized)
		return
	}
	sessionID := cookie.Value

	vendorID, err := a.DB.VendorCheckSession(sessionID)
	if err != nil || vendorID == "" {
		http.Error(w, "Invalid or expired session", http.StatusUnauthorized)
		fmt.Println("error:", err, "vendorID:", vendorID)
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

	//POST Method (VendorSetDiscount) -- INSERT the data back to DB once vendor has set which meal to be live for discount and the quantity
	//POST Method (VendorSetLaunch) -- UPDATE this is for updating the DiscountStart and DiscountEnd time for launch/end of discount
	if r.Method == http.MethodPost {
		defer r.Body.Close()

		var vendorDiscount models.VendorLaunch

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		fmt.Println("print req body", r.Body)
		err := json.NewDecoder(r.Body).Decode(&vendorDiscount)

		if err != nil {
			http.Error(w, "Failed to decode vendor discount details", http.StatusBadRequest)
			return
		}

		const timeLayout = "15:04"
		if _, err := time.Parse(timeLayout, vendorDiscount.DiscountStart); err != nil {
			log.Fatalf("Invalid DiscountStart format: %v", err)
		}
		if _, err := time.Parse(timeLayout, vendorDiscount.DiscountEnd); err != nil {
			log.Fatalf("Invalid DiscountEnd format: %v", err)
		}

		//fmt.Println("this is time", fullDateTimeStart, fullDateTimeEnd)
		//if vendorDiscount.DiscountStart == "00:00" {
		//	vendorDiscount.DiscountStart = "0001-01-01 00:00:00"
		//} else {
		//	vendorDiscount.DiscountStart = slices.Concat(time.Now(), vendorDiscount.DiscountStart)
		//}
		//
		//if vendorDiscount.DiscountEnd == "00:00" {
		//	vendorDiscount.DiscountEnd = "0001-01-01 00:00:00"
		//}

		const defaultDate = "0001-01-01"
		fmt.Println("this is initial", vendorDiscount.DiscountStart, vendorDiscount.DiscountEnd)

		fullDateTimeStart := defaultDate + " " + vendorDiscount.DiscountStart + ":00"
		fullDateTimeEnd := defaultDate + " " + vendorDiscount.DiscountEnd + ":00"
		fmt.Println("This is time:", fullDateTimeStart, fullDateTimeEnd)

		const dateTimeLayout = "2006-01-02 15:04:05"
		start, err := time.Parse(dateTimeLayout, fullDateTimeStart)
		if err != nil {
			http.Error(w, "Invalid DiscountStart format", http.StatusBadRequest)
			return
		}
		vendorDiscount.DiscountStart = start.Format(dateTimeLayout)

		end, err := time.Parse(dateTimeLayout, fullDateTimeEnd)
		if err != nil {
			http.Error(w, "Invalid DiscountEnd format", http.StatusBadRequest)
			return
		}
		vendorDiscount.DiscountEnd = end.Format(dateTimeLayout)

		// Create Frontend struct
		frontend := Frontend{
			button: vendorDiscount.Button,
			start:  start,
			end:    end,
		}

		fmt.Println("D BUTTON", frontend.button, vendorDiscount.Button)

		// Handle the schedule button
		vendor := &models.Vendor{} // Assume an existing Vendor struct
		vendor = handleScheduleButton(frontend, vendor)
		//fmt.Println("vendor is discount", vendorDiscount.IsDiscountOpen)
		vendorDiscount.IsDiscountOpen = vendor.IsDiscountOpen
		postResults, err := a.DB.VendorSetDiscount(&vendorDiscount)
		if err != nil {
			http.Error(w, "Failed to process discounts", http.StatusInternalServerError)
			return
		}

		// Respond back to the client
		w.Header().Set("Content-Type", "application/json")
		response := map[string]bool{
			"success": postResults,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (a *Apiserver) GetCustomerDiscount(w http.ResponseWriter, r *http.Request) {
	// get list of vendors
	/* select * from Vendor where IsDiscountOpen = true OR (current_time() >= DiscountStart AND current_time() <= DiscountEnd); */

	query := fmt.Sprintf("SELECT * FROM Vendor WHERE IsDiscountOpen = true OR (CURRENT_TIME() >= DiscountStart AND CURRENT_TIME() <= DiscountEnd)")
	rows, err := a.DB.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch vendor discount details", http.StatusUnauthorized)
		fmt.Println("Error fetching vendor discount details: ", err)
		return
	}
	defer rows.Close()

	fmt.Println(rows)
	var result []models.Vendor
	for rows.Next() {
		var vendor models.Vendor
		if err := rows.Scan(
			&vendor.VendorID,
			&vendor.VendorName,
			&vendor.Address,
			&vendor.IsOpen,
			&vendor.IsDiscountOpen,
			&vendor.DiscountStart,
			&vendor.DiscountEnd,
			&vendor.Password,
			&vendor.VendorImage); err != nil {
			http.Error(w, "Failed to fetch vendor discount details", http.StatusUnauthorized)
			fmt.Println("Error fetching vendor discount details: ", err)
			return
		}
		result = append(result, vendor)
	}
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
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
		discountedMeals, dberr := a.DB.GetDiscountedMealsFromVendor(meals) //[]models.Discount
		if dberr != nil {
			fmt.Println(dberr)
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "No discounted meals found"}`))
			return
		}

		//fmt.Println(discountedMealsAndMealName)
		vendorName, vendorerr := a.DB.FetchVendorName(vendorID)
		if vendorerr != nil {
			fmt.Println(vendorerr)
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "No vendor found"}`))
			return
		}

		discountedMealsAndMealName, dberr2 := a.DB.MapMealIDAndMealName(discountedMeals, vendorName) //DiscountAndMealName
		if dberr2 != nil {
			fmt.Println(dberr2)
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "No discounted meals found and mapping"}`))
			return
		}

		jsonerr := json.NewEncoder(res).Encode(discountedMealsAndMealName) //json data returned
		if jsonerr != nil {
			fmt.Println(jsonerr)
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte(`{"error": "Error with json conversion. Discounted meals not found"}`))
			return
		}
	}
}

func (a *Apiserver) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}
