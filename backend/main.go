package main

import (
	"fmt"
	"hackathon/database"
	"hackathon/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// initiate new db instance
	db := database.NewDBClient()

	//initialize api server
	apiServer := server.NewApiserver(db)
	router := mux.NewRouter() // main router for your app
	handler := enableCORS(router)
	apiServer.RegisterRoutes(router)

	// create server
	fmt.Println("Listening to port 5001")
	log.Fatal(http.ListenAndServe(":5001", handler))
}
