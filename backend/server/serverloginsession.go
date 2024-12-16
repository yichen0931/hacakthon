package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const SESSION_EXPIRY_TIME = 24

func (a *Apiserver) Login(res http.ResponseWriter, req *http.Request) {
	if http.MethodPost == req.Method {
		if req.Header.Get("Content-Type") == "application/json" {
			// Create a map to hold the decoded data
			res.Header().Set("Content-Type", "application/json")
			loginDetails := make(map[string]string) //vendorID or customerID, password and Role

			// Decode the JSON from the request body into the map
			err := json.NewDecoder(req.Body).Decode(&loginDetails)
			if err != nil {
				fmt.Println("Error with JSON decoding of req body in login:", err)
				return
			}

			//fmt.Println("Username:", loginDetails["Username"])
			//fmt.Println("Password:", loginDetails["Password"])
			//fmt.Println("Vendor:", loginDetails["Role"])

			getRole := loginDetails["Role"]
			switch getRole {
			case "Vendor":
				vendorID := loginDetails["UserID"]
				password := loginDetails["Password"]
				loginerr := a.DB.VendorLogin(vendorID, password)
				if loginerr != nil {
					res.WriteHeader(http.StatusNotFound)
					//res.Write([]byte(`{"error": "Login error"}`))
					res.Write([]byte(`{"authenticated": false, "UserID": ""}`))
					return
				}

				//insert into session
				sessionID, sessionErr := a.DB.VendorAddSession(vendorID) //add vendor into session
				if sessionErr != nil {
					res.WriteHeader(http.StatusNotFound)
					//res.Write([]byte(`{"error": "session error"}`))
					res.Write([]byte(`{"authenticated": false, "UserID": ""}`))
					return
				}

				//create session in client side
				sessionCookie := &http.Cookie{ //create & populate the session model
					Name:    "vendorSessionCookie",
					Value:   sessionID,
					Expires: time.Now().Add(SESSION_EXPIRY_TIME * time.Hour), //Changeable. NOTE, browser will auto delete this once expires client side.
					Path:    "/",
					Domain:  "localhost", //only this domain
				}
				http.SetCookie(res, sessionCookie)
				res.Write([]byte(`{"authenticated": true, "Role":"Vendor" , "UserID": "` + vendorID + `"}`))
				//http.Redirect(res, req, "/vendor/discount", http.StatusSeeOther)
			case "Customer":
				customerID := loginDetails["UserID"]
				password := loginDetails["Password"]
				loginerr := a.DB.CustomerLogin(customerID, password)
				if loginerr != nil {
					res.WriteHeader(http.StatusNotFound)
					//res.Write([]byte(`{"error": "Login error"}`))
					res.Write([]byte(`{"authenticated": false, "UserID": ""}`))
					return
				}

				//insert into session
				sessionID, sessionErr := a.DB.CustomerAddSession(customerID) //add customer into session
				if sessionErr != nil {
					res.WriteHeader(http.StatusNotFound)
					//res.Write([]byte(`{"error": "session error"}`))
					res.Write([]byte(`{"authenticated": false, "UserID": ""}`))
					return
				}

				//create session in client side
				sessionCookie := &http.Cookie{ //create & populate the session model
					Name:    "customerSessionCookie",
					Value:   sessionID,
					Expires: time.Now().Add(SESSION_EXPIRY_TIME * time.Hour), //Changeable. NOTE, browser will auto delete this once expires client side.
					Path:    "/",
					Domain:  "localhost", //only this domain
				}
				http.SetCookie(res, sessionCookie)
				res.Write([]byte(`{"authenticated": true, "Role":"Customer", "UserID": "` + customerID + `"}`))
				//http.Redirect(res, req, "/customer/discount", http.StatusSeeOther)
			default:
				fmt.Println("Invalid role type")
				res.WriteHeader(http.StatusNotFound)
				//res.Write([]byte(`{"error": "Role doesnt exist"}`))
				res.Write([]byte(`{"authenticated": false, "UserID": ""}`))
				return
			}
		}

	}
}
