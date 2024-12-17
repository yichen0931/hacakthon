package models

import "time"

type VendorView struct {
	IsOpen                    bool   `json:"isOpen"`
	IsDiscount                bool   `json:"isDiscount"`
	DiscountStart             string `json:"discountStart"`
	DiscountEnd               string `json:"discountEnd"`
	MealID                    string `json:"mealID"`
	MealName                  string `json:"mealName"`
	Description               string `json:"description"`
	Availability              bool   `json:"availability"`
	SustainabilityCreditScore int    `json:"sustainabilityCreditScore"`
}

type VendorLaunch struct {
	Discount       []Discount `json:"Meals"`
	DiscountStart  string     `json:"DiscountStart"`
	StartTime      time.Time
	DiscountEnd    string `json:"DiscountEnd"`
	EndTime        time.Time
	Button         string `json:"Button"`
	IsDiscountOpen bool
}

// mus added
type DiscountAndMealName struct {
	MealID        string  `json:"MealID"`
	DiscountPrice float64 `json:"DiscountPrice"`
	Quantity      int     `json:"Quantity"`
	MealName      string  `json:"MealName"`
	MealPrice     float64 `json:"MealPrice"`
	VendorName    string  `json:"VendorName"`
}
