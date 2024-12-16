package database

import (
	"errors"
	"fmt"
	"hackathon/models"
)

func (db *DBClient) InsertOrder(order models.Orders) error {
	dbString := fmt.Sprintf("INSERT INTO Orders VALUES ('%s','%s','%s','%s','%s','%v','%s')",
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
		dbString := fmt.Sprintf("INSERT INTO OrderDetail VALUES ('%s','%s','%v','%v')",
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
