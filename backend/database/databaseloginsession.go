package database

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const SESSION_EXPIRY_TIME = 24

func (db *DBClient) VendorLogin(vendorID, inputPassword string) error {
	dbString := fmt.Sprintf("SELECT VendorID FROM Vendor WHERE VendorID='%s' AND Password='%s'", vendorID, inputPassword)
	result, err := db.DB.Query(dbString)
	if err != nil {
		fmt.Println("Error querying database")
		return err
	}

	if result.Next() {
		fmt.Println("Vendor ID and password matches! Login successful...creating session")
		return nil
	} else { //no such login
		return errors.New("No such vendor!")
	}
}

func (db *DBClient) VendorAddSession(vendorID string) (string, error) {
	sessionExpiryTime := time.Now().Add(SESSION_EXPIRY_TIME * time.Hour) //this session expiry is changeable
	expiryString := sessionExpiryTime.Format("2006-01-02 15:04:05")      //layout matching the format of the string
	sessionID := uuid.NewString()

	sessionString := fmt.Sprintf("INSERT INTO VendorSessions(SessionID, VendorID, SessionExpiry) VALUES ('%s','%s','%s')",
		sessionID, vendorID, expiryString)

	result, err := db.DB.Exec(sessionString)
	if err != nil {
		return "", err
	}

	if val, _ := result.RowsAffected(); val != 0 {
		fmt.Println("Successfully added session") //return session ID
		return sessionID, nil
	} else {
		fmt.Println("Failed to add session for vendor")
		return "", errors.New("Failed to add session for vendor")
	}
}

func (db *DBClient) CustomerLogin(customerID, inputPassword string) error {
	dbString := fmt.Sprintf("SELECT CustomerID FROM Customer WHERE CustomerID='%s' AND Password='%s'", customerID, inputPassword)
	result, err := db.DB.Query(dbString)
	if err != nil {
		fmt.Println("Error querying database")
		return err
	}

	if result.Next() {
		fmt.Println("Customer ID and password matches! Login successful...creating session")
		return nil
	} else { //no such login
		return errors.New("No such customer!")
	}
}

func (db *DBClient) CustomerAddSession(customerID string) (string, error) {
	sessionExpiryTime := time.Now().Add(SESSION_EXPIRY_TIME * time.Hour) //this session expiry is changeable
	expiryString := sessionExpiryTime.Format("2006-01-02 15:04:05")      //layout matching the format of the string
	sessionID := uuid.NewString()

	sessionString := fmt.Sprintf("INSERT INTO CustomerSessions(SessionID, CustomerID, SessionExpiry) VALUES ('%s','%s','%s')",
		sessionID, customerID, expiryString)

	result, err := db.DB.Exec(sessionString)
	if err != nil {
		return "", err
	}

	if val, _ := result.RowsAffected(); val != 0 {
		fmt.Println("Successfully added session") //return session ID
		return sessionID, nil
	} else {
		fmt.Println("Failed to add session for customer")
		return "", errors.New("Failed to add session for customer")
	}
}

func (db *DBClient) CheckSessionExistVendor(vendorSessionCookie string) string {
	dbString := fmt.Sprintf("SELECT VendorID FROM VendorSessions WHERE SessionID='%s'", vendorSessionCookie)
	result, err := db.DB.Query(dbString)
	if err != nil {
		fmt.Println("Error querying database")
		return ""
	}

	var vendorID string
	if result.Next() {
		result.Scan(&vendorID)
		return vendorID
	} else {
		return ""
	}
}

func (db *DBClient) CheckSessionExistCustomer(customerSessionCookie string) string {
	dbString := fmt.Sprintf("SELECT CustomerID FROM CustomerSessions WHERE SessionID='%s'", customerSessionCookie)
	result, err := db.DB.Query(dbString)
	if err != nil {
		fmt.Println("Error querying database")
		return ""
	}

	var customerID string
	if result.Next() {
		result.Scan(&customerID)
		return customerID
	} else {
		return ""
	}
}
