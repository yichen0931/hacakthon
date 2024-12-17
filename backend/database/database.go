package database

import (
	"database/sql"
	"errors"
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

	//query := fmt.Sprintf("SELECT v.IsOpen, v.IsDiscountOpen AS IsDiscount, v.DiscountStart, v.DiscountEnd, , m.MealID, m.MealName, m.Description, m.Availability, m.SustainabilityCreditScore FROM Vendor v LEFT JOIN Meal m ON v.VendorID = m.VendorID WHERE v.VendorID = '%s'", vendorID)

	//this is only for testing purpose (need to make sure that DiscountStart, DiscountEnd is NOT NULL)
	query := fmt.Sprintf("SELECT v.IsOpen, v.IsDiscountOpen AS IsDiscount, v.DiscountStart, v.DiscountEnd, m.MealID, m.MealName, m.Description, m.Availability, m.SustainabilityCreditScore, m.Price FROM Vendor v LEFT JOIN Meal m ON v.VendorID = m.VendorID WHERE v.VendorID = '%s'", vendorID)

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal("Failed to get data from database", err.Error())
		return nil, err
	}

	defer rows.Close()
	//var discountStart sql.NullTime
	//var discountEnd sql.NullTime
	for rows.Next() {
		fmt.Println("inside")
		var vendorView models.VendorView
		fmt.Println("scanning")
		err := rows.Scan(
			&vendorView.IsOpen,
			&vendorView.IsDiscount,
			&vendorView.DiscountStart,
			&vendorView.DiscountEnd,
			//&discountStart,
			//&discountEnd,
			&vendorView.MealID,
			&vendorView.MealName,
			&vendorView.Description,
			&vendorView.Availability,
			&vendorView.SustainabilityCreditScore,
			&vendorView.MealPrice,
		)

		//if !discountStart.Valid {
		//	vendorView.DiscountStart = "0001-01-01 00:00:00"
		//}
		//
		//if !discountEnd.Valid {
		//	vendorView.DiscountEnd = "0001-01-01 00:00:00"
		//}
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
// Updating the discount start time, end time which are ready for launch on the Vendor side
func (db *DBClient) VendorSetDiscount(vendorLaunch *models.VendorLaunch, vendorID string) (bool, error) {
	// Iterate over each discount
	for _, discount := range vendorLaunch.Discount {
		query := fmt.Sprintf("UPDATE Discount SET DiscountedPrice = %v, Quantity = %d WHERE MealID ='%s'", discount.DiscountPrice, discount.Quantity, discount.MealID)
		fmt.Println(query)
		_, err := db.DB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to insert Discounted meal item: %s", err.Error())
			return false, err
		}
		fmt.Printf("Successfully inserted discount for MealID %s\n", discount.MealID)
	}

	// Build the update query based on whether the times were parsed successfully
	//timeValue := time.Now()
	updateQuery := fmt.Sprintf("UPDATE Vendor SET IsDiscountOpen = %t, DiscountStart = '%s', DiscountEnd = '%s' WHERE VendorID = '%s'", vendorLaunch.IsDiscountOpen, vendorLaunch.DiscountStart, vendorLaunch.DiscountEnd, vendorID)
	fmt.Println("update query", updateQuery)

	// Execute the query with the appropriate parameters
	_, err := db.DB.Exec(updateQuery)
	if err != nil {
		log.Fatalf("Failed to update Vendor Discount status: %s", err.Error())
		return false, err
	}

	return true, nil
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
		dberr := result.Scan(&meal.MealID, &meal.VendorID, &meal.MealName, &meal.Description, &meal.Price, &meal.Availability, &meal.SustainabilityCreditScore, &meal.VendorImage)
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

func (db *DBClient) MapMealIDAndMealName(discounts []models.Discount, vendorName string) ([]models.DiscountAndMealName, error) {
	var finalResult []models.DiscountAndMealName
	builder := models.DiscountAndMealName{}
	for _, discount := range discounts {
		builder.MealID = discount.MealID
		builder.DiscountPrice = discount.DiscountPrice
		builder.Quantity = discount.Quantity

		//fetch meal name via meal id
		dbString := fmt.Sprintf("SELECT MealName, Price FROM Meal WHERE MealID='%s'", discount.MealID)
		result, err := db.DB.Query(dbString)
		if err != nil {
			fmt.Println("Error in MapMealIDAndMealName", err)
			return nil, err
		}

		var mealname string
		var mealprice float64
		if result.Next() {
			result.Scan(&mealname, &mealprice)
			builder.MealName = mealname
			builder.MealPrice = mealprice
			builder.VendorName = vendorName
			finalResult = append(finalResult, builder)
		} else {
			continue
		}
	}
	return finalResult, nil
}

func (db *DBClient) FetchVendorName(vendorID string) (string, error) {
	dbString := fmt.Sprintf("SELECT VendorName FROM Vendor WHERE VendorID='%s'", vendorID)
	result, err := db.DB.Query(dbString)
	if err != nil {
		return "", err
	}

	var vendorName string
	if result.Next() {
		result.Scan(&vendorName)
		return vendorName, nil
	} else {
		return "", errors.New("Vendor name not found")
	}
}
