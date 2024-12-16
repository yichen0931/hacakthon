package database

import (
	"database/sql"
	"fmt"
	"hackathon/models"
	"log"
)

type DBClient struct {
	DB *sql.DB
}

func NewDBClient() *DBClient {
	dsn := "user:localhost:strongpassword@tcp(localhost:3306)/dealsDB"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	//ping our database to check if the credentials are valid
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database %v", err)
	}

	fmt.Println("Successfully connected to MySQL database") // check our connection
	newDB := &DBClient{DB: db}
	return newDB
}

func (db *DBClient) DiscountStatus(res string) error {
	status := models.Vendor{}
	fmt.Println("inside", res)
	if res == "Launch" {
		status.IsDiscountOpen = true
		return nil
	} else if res == "End" {
		status.IsDiscountOpen = false
		return nil
	}
	return fmt.Errorf("Invalid status %v", res)
}
