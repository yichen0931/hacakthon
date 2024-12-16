package database

import (
	"database/sql"
	"fmt"
	"hackathon/models"
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

	// ping our database to check if the credentials are valid
	//	if err := db.Ping(); err != nil {
	//		log.Fatalf("Failed to ping database %v", err)
	//	}

	fmt.Println("Successfully connected to MySQL database") // check our connection
	newDB := &DBClient{DB: db}
	return newDB
}

//func (db *DBClient) VendorGetMeals() {
//	results, err := db.DB.Query("SELECT * FROM Meal")
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	defer results.Close()
//
//	for results.Next() {
//		var meal models.Meal
//		err = results.Scan(
//			&meal.MealID,
//			&meal.MealName,
//			&meal.Description,
//			&meal.Price)
//
//		if err != nil {
//			panic(err.Error())
//		}
//
//		fmt.Printf("Meals successfully retrieved %s %s %s %d\n", meal.MealID, meal.MealName, meal.Description, meal.Price)
//	}
//}

//func (db *DBClient) SetDiscountTime(startTime string, endTime string) {
//	query := fmt.Sprintf("UPDATE Vendor SET DiscountStart = '%s', DiscountEnd = '%s'", startTime, endTime)
//
//	if _, err := db.DB.Exec(query); err != nil {
//		log.Fatalf("Failed to update Vendor Discount time %v", err)
//		return
//	}
//	log.Printf("Successfully updated Vendor Discount time to start at '%s', and end at '%s'", startTime, endTime)
//}

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

//func (db *DBClient) SetDiscountedPrice(status *models.Vendor, meal *models.Meal) {
//	if status.IsDiscountOpen == true {
//
//	} else if status.IsDiscountOpen == false || status.DiscountStart != time.Now().Format("2006-01-02 15:04:05") {
//		log.Fatalln("Discount ")
//	}
//}
