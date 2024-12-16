package database

import (
	"database/sql"
	"fmt"
	"log"
)

type DBClient struct {
	DB *sql.DB
}

func NewDBClient() *DBClient {
	dsn := "user:strongpassword@tcp(127.0.0.1:3306)/dealsDB"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close() // defer the close
	
	// ping our database to check if the credentials are valid
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database %v", err)
	}

	fmt.Println("Successfully connected to MySQL database") // check our connection

	newDB := &DBClient{DB: db}
	return newDB
}
