package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"hackathon/models"
	"net/http"
)

type Apiserver struct {
	DB    *sql.DB
	Foods map[string]models.FoodInfo
	Users map[string]models.User
}

func NewApiserver(db *sql.DB) *Apiserver {
	return &Apiserver{
		DB:    db,
		Foods: make(map[string]models.FoodInfo),
		Users: make(map[string]models.User),
	}
}

func (a *Apiserver) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/panda/v1/", a.Home) // handler "home"
	router.HandleFunc("/panda/v1/food", a.Allfood)
	router.HandleFunc("/panda/v1/food/{foodid}", a.Food).Methods("GET", "POST", "PUT")
	router.HandleFunc("/panda/v1/users", a.GetAllUsers).Methods("GET")
	router.HandleFunc("/panda/v1/adduser", a.AddUser).Methods("POST")
}

func (a *Apiserver) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	results, err := a.DB.Query("SELECT * FROM user")

	if err != nil {
		panic(err.Error())
	}

	defer results.Close()

	// iterate all people
	var users []models.User
	for results.Next() {
		var user models.User
		err = results.Scan(
			&user.UserID,
			&user.Username,
			&user.UserPassword,
			&user.Firstname,
			&user.Lastname,
			&user.UserRole,
			&user.Salt,
		)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
		fmt.Printf("Record found: %s: %s (%s %s) \n", user.UserID, user.Username, user.Firstname, user.Lastname)
	}
	json.NewEncoder(w).Encode(users)
}

func (a *Apiserver) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user.UserID = "12345"
	user.UserRole = "human"
	user.Salt = user.UserID + user.UserPassword
	fmt.Println(user)
	query := fmt.Sprintf("INSERT INTO user(userID, username, userpassword, firstname, lastname, userrole, salt) VALUES('%v', '%v', '%v', '%v', '%v', '%v', '%v')", user.UserID, user.Username, user.UserPassword, user.Firstname, user.Lastname, user.UserRole, user.Salt)

	_, err := a.DB.Exec(query)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("User Record Inserted Successfully")
	json.NewEncoder(w).Encode(user)
}

func (a *Apiserver) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}

func (a *Apiserver) Allfood(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list of all the food")

	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v)
	}

	// return nicely in JSON format
	json.NewEncoder(w).Encode(a.Foods)
}

func (a *Apiserver) Food(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "detail..."+params["foodid"])

	// do some query with the parameter result

	// or process it
}
