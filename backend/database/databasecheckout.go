package database

import (
	"errors"
	"fmt"
	"hackathon/models"
)

func (db *DBClient) InsertOrder(order models.Orders) error {
	dbString := fmt.Sprintf("INSERT INTO Orders VALUES ('%s','%s','%s','%s','%s',%v,'%s')",
		order.OrderID, order.CustomerID, order.RiderID, order.OrderStatus, order.OrderEnd, order.Total, order.DeliveryAddress)

	result, err := db.DB.Exec(dbString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if val, _ := result.RowsAffected(); val != 0 {
		fmt.Println("Order inserted", order.OrderID)
		return nil
	} else {
		fmt.Println("order is not inserted")
		return errors.New("order is not inserted")
	}
}

func (db *DBClient) InsertOrderDetail(orderDetails []models.OrderDetail) {
	for _, orderDetail := range orderDetails {
		dbString := fmt.Sprintf("INSERT INTO OrderDetail VALUES ('%s','%s',%v,%v)",
			orderDetail.OrderID, orderDetail.MealID, orderDetail.MealQty, orderDetail.MealPrice)
		fmt.Println(dbString)

		result, err := db.DB.Exec(dbString)
		if err != nil {
			fmt.Println(err)
			return
		}

		if val, _ := result.RowsAffected(); val != 0 {
			fmt.Println("Order inserted", orderDetail.OrderID)
		} else {
			fmt.Println("order is not inserted")
		}
	}
}

func (db *DBClient) ValidateDiscountStockQty(meals []models.OrderDetail) bool {
	//for each meal ordered, check stock in discounts table
	for _, meal := range meals { //if any exceeds, then return false
		mealID := meal.MealID
		dbString := fmt.Sprintf("SELECT Quantity FROM Discount WHERE MealID='%s'", mealID)
		result, err := db.DB.Query(dbString)
		if err != nil {
			fmt.Println(err)
			return false
		}

		var tableQty int
		if result.Next() {
			//do the checking of qty here.
			result.Scan(&tableQty)
			if meal.MealQty > tableQty { //my wanted qty exceeds
				return false
			} else { //mealqty = tableqty (0=0 ok) or 2 < 3
				continue //check next
			}
		} else { //empty
			return false
		}

	}
	return true //this means all qty is ok so customer can buy!
}

// reduce qty in discounts table
func (db *DBClient) ReduceQty(meals []models.OrderDetail) error {
	for _, meal := range meals {
		mealID := meal.MealID

		//first select from db the qty of that meal
		dbString := fmt.Sprintf("SELECT Quantity FROM Discount WHERE MealID='%s'", mealID)
		result, err := db.DB.Query(dbString)
		if err != nil {
			fmt.Println(err)
			return err
		}

		var tableQty int
		if result.Next() {
			result.Scan(&tableQty)
			resultingQty := tableQty - meal.MealQty
			if resultingQty < 0 { //if negative
				fmt.Println("Not enough quantity to reduce!!!")
				return errors.New("Not enough quantity to reduce")
			}
			//update discount table
			updateString := fmt.Sprintf("UPDATE Discount SET Quantity=%d WHERE MealID='%s'", resultingQty, mealID)
			_, dberr := db.DB.Exec(updateString)
			if dberr != nil {
				fmt.Printf("something wrong with updating qty [%s] for mealid %s \n", dberr, meal.MealID)
				return dberr
			}
			fmt.Printf("Successfully reduced Qty of mealid: [%s] from %v to %v (-%v)\n", meal.MealID, tableQty, resultingQty, meal.MealQty)

		} else { //empty
			return errors.New("meal not in discount")
		}
	}
	return nil
}
