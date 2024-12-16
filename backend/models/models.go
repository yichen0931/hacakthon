package models

type Vendor struct {
	VendorID       string
	VendorName     string
	Address        string
	IsOpen         bool
	IsDiscountOpen bool
	DiscountStart  int64
	DiscountEnd    int64
}

type Meal struct {
	MealID                    string
	VendorID                  string
	MealName                  string
	Description               string
	Price                     int64
	Availability              bool
	SustainabilityCreditScore int
}

type Rider struct {
	RiderID      string
	RiderName    string
	VehiclePlate string
	Availability bool
}

type Customer struct {
	CustomerID                           string
	CustomerName                         string
	Address                              string
	AccumulatedSustainabilityCreditScore int
}

type Discount struct {
	MealID        string
	DiscountPRice int64
	Quantity      int
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
	OrderID         string
	CustomerID      string
	RiderID         string
	OrderStatus     OrderStatus
	OrderEnd        int64
	Total           int64
	DeliveryAddress string
}

type OrderDetail struct {
	OrderID   string
	MealID    string
	MealQty   int
	MealPrice int64
}

type Sessions struct {
	SessionID     string
	UserID        string
	SessionExpiry int64
}

type Role int

const (
	VENDOR Role = iota
	CUSTOMER
	RIDER
)

type Users struct {
	UserID   string
	UserName string
	Password string
	Role     Role
}
