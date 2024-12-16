package models

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
