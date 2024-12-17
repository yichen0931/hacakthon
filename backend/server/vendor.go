package server

import (
	"fmt"
	"hackathon/models"
	"time"
)

func ConvertStringToTime(t string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, t)
	return parsedTime, err
}

func ConvertTimeToString(t time.Time) string {
	timeString := time.Now().Format("2006-01-02 15:04:05")
	return timeString
}

type Frontend struct {
	button string
	start  time.Time
	end    time.Time
}

func handleScheduleButton(frontend Frontend, vendor *models.Vendor) *models.Vendor {
	fmt.Println("lawl", frontend.button, vendor)
	switch frontend.button {
	case "Launch":
		fmt.Println("it launcehd", vendor.IsDiscountOpen)
		vendor.IsDiscountOpen = true
		fmt.Println("it checkue", vendor.IsDiscountOpen)

	case "Schedule":
		t := time.Now()
		if t.After(frontend.start) && t.Before(frontend.end) {
			vendor.IsDiscountOpen = true
		} else {
			vendor.IsDiscountOpen = false
		}

	case "End":
		vendor.IsDiscountOpen = false

	default:
		fmt.Println("invalid button string")
	}
	vendor.DiscountStart = ConvertTimeToString(frontend.start)
	vendor.DiscountEnd = ConvertTimeToString(frontend.end)
	fmt.Println(vendor)
	return vendor
}
