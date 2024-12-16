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
	dsn := "root:strongpassword@tcp(localhost:3307)/dealsDB"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	// ping our database to check if the credentials are valid
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
func (db *DBClient) GetMealFromVendor(vendorID string) ([]models.Meal, error) {
	var meals []models.Meal

	dbString := fmt.Sprintf("SELECT * FROM Meal WHERE VendorID = '%s'", vendorID)
	result, err := db.DB.Query(dbString)
	if err != nil {
		fmt.Println("Error in GetMealFromVendor", err)
		return nil, err
	}
	defer result.Close()

	meal := models.Meal{}
	for result.Next() {
		dberr := result.Scan(&meal.MealID, &meal.VendorID, &meal.MealName, &meal.Description, &meal.Price, &meal.Availability, &meal.SustainabilityCreditScore)
		if dberr != nil {
			fmt.Println(dberr)
			return nil, dberr
		}
		meals = append(meals, meal)
	}

	return meals, nil
}

func (db *DBClient) GetDiscountedMealsFromVendor(meals []models.Meal) ([]models.Discount, error) {
	var discountedMeals []models.Discount
	for _, meal := range meals {
		dbString := fmt.Sprintf("SELECT * FROM Discount WHERE MealID='%s'", meal.MealID)
		result, err := db.DB.Query(dbString)
		if err != nil {
			fmt.Println("Error in GetDiscountedMealsFromVendor", err)
			return nil, err
		}
		var discountedMeal = models.Discount{}
		if result.Next() {
			result.Scan(&discountedMeal.MealID, &discountedMeal.DiscountPrice, &discountedMeal.Quantity)
			discountedMeals = append(discountedMeals, discountedMeal)
		} else {
			continue
			//return nil, errors.New("empty rows")
		}
	}
	return discountedMeals, nil
}
