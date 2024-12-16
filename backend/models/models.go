package models

type Vendor struct {
	VendorID       string
	VendorName     string
	Address        string
	IsOpen         bool
	IsDiscountOpen bool
	DiscountStart  string
	DiscountEnd    string
	Password       string
}

type Meal struct {
	MealID                    string
	VendorID                  string
	MealName                  string
	Description               string
	Price                     float64
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
	Password                             string
}

type Discount struct {
	MealID        string
	DiscountPrice float64
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
	OrderEnd        string
	Total           float64
	DeliveryAddress string
}

type OrderDetail struct {
	OrderID   string
	MealID    string
	MealQty   int
	MealPrice float64
}

type CustomerSessions struct {
	SessionID     string
	CustomerID    string
	SessionExpiry string
}

type VendorSessions struct {
	SessionID     string
	VendorID      string
	SessionExpiry string
}
