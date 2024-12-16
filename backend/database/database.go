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

// View the (home) page for Vendor once they are logged in
func (db *DBClient) VendorViewAllMeal(vendorID string) ([]models.VendorView, error) {
	var vendorViews []models.VendorView

	//query := fmt.Sprintf("SELECT v.IsOpen, v.IsDiscountOpen AS IsDiscount, v.DiscountStart, v.DiscountEnd, m.MealID, m.MealName, m.Description, m.Availability, m.SustainabilityCreditScore FROM Vendor v LEFT JOIN Meal m ON v.VendorID = m.VendorID WHERE v.VendorID = '%s'", vendorID)

	//this is only for testing purpose (need to make sure that DiscountStart, DiscountEnd is NOT NULL)
	query := fmt.Sprintf("SELECT v.IsOpen, v.IsDiscountOpen AS IsDiscount, v.DiscountStart, v.DiscountEnd, m.MealID, m.MealName, m.Description, m.Availability, m.SustainabilityCreditScore FROM Vendor v LEFT JOIN Meal m ON v.VendorID = m.VendorID WHERE v.VendorID = 'V002'")

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal("Failed to get data from database", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		fmt.Println("inside")
		var vendorView models.VendorView
		fmt.Println("scanning")
		err := rows.Scan(
			&vendorView.IsOpen,
			&vendorView.IsDiscount,
			&vendorView.DiscountStart,
			&vendorView.DiscountEnd,
			&vendorView.MealID,
			&vendorView.MealName,
			&vendorView.Description,
			&vendorView.Availability,
			&vendorView.SustainabilityCreditScore,
		)
		fmt.Println("vendor view", vendorView)
		if err != nil {
			log.Fatalln("Failed to scan row for vendor views", err.Error())
			return nil, err
		}
		vendorViews = append(vendorViews, vendorView)
		fmt.Println(vendorViews)
	}
	return vendorViews, nil
}

// Inserting the meals discounts/prices and quantity on Vendor side
func (db *DBClient) VendorSetDiscount(updatedDiscount *models.VendorSetDiscount) error {

	query := fmt.Sprintf("INSERT INTO Discount (MealID, DiscountedPrice, Quantity) VALUES ('%s', '%.2f', '%d')", updatedDiscount.MealID, updatedDiscount.DiscountedPrice, updatedDiscount.Quantity)

	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to insert Discounted meal item: %s", err.Error())
		return err
	}

	fmt.Println("Successfully inserted Discounted meal item")
	return nil

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
