package models

type Vendor struct {
	VendorID       string `json:"VendorID"`
	VendorName     string `json:"VendorName"`
	Address        string `json:"Address"`
	IsOpen         bool   `json:"IsOpen"`
	IsDiscountOpen bool   `json:"IsDiscountOpen"`
	DiscountStart  int64  `json:"DiscountStart"`
	DiscountEnd    int64  `json:"DiscountEnd"`
}

type Meal struct {
	MealID                    string `json:"MealID"`
	VendorID                  string `json:"VendorID"`
	MealName                  string `json:"MealName"`
	Description               string `json:"Description"`
	Price                     int64  `json:"Price"`
	Availability              bool   `json:"Availability"`
	SustainabilityCreditScore int    `json:"SustainabilityCreditScore"`
}

type Rider struct {
	RiderID      string `json:"RiderID"`
	RiderName    string `json:"RiderName"`
	VehiclePlate string `json:"VehiclePlate"`
	Availability bool   `json:"Availability"`
}

type Customer struct {
	CustomerID                           string `json:"CustomerID"`
	CustomerName                         string `json:"CustomerName"`
	Address                              string `json:"Address"`
	AccumulatedSustainabilityCreditScore int    `json:"AccumulatedSustainabilityCreditScore"`
}

type Discount struct {
	MealID        string `json:"MealID"`
	DiscountPrice int64  `json:"DiscountPrice"`
	Quantity      int    `json:"Quantity"`
}

type OrderStatus int

const (
	CART OrderStatus = iota
	PENDING
	ORDERRECEIVED
	GROUPORDER
	PREPARING
	PICKED
	DELIVERED
)

type Orders struct {
	OrderID         string      `json:"OrderID"`
	CustomerID      string      `json:"CustomerID"`
	RiderID         string      `json:"RiderID"`
	OrderStatus     OrderStatus `json:"OrderStatus"`
	OrderEnd        int64       `json:"OrderEnd"`
	Total           int64       `json:"Total"`
	DeliveryAddress string      `json:"DeliveryAddress"`
}

type OrderDetail struct {
	OrderID   string `json:"OrderID"`
	MealID    string `json:"MealID"`
	MealQty   int    `json:"MealQty"`
	MealPrice int64  `json:"MealPrice"`
}

type Sessions struct {
	SessionID     string `json:"SessionID"`
	UserID        string `json:"UserID"`
	SessionExpiry int64  `json:"SessionExpiry"`
}

type Role int

const (
	VENDOR Role = iota
	CUSTOMER
	RIDER
)

type Users struct {
	UserID   string `json:"UserID"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Role     Role   `json:"Role"`
}
