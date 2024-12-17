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
	timeString := time.Now().Add(1 * time.Hour).Format("2006-01-02 15:04:05")
	return timeString
}

type Frontend struct {
	button string
	start  time.Time
	end    time.Time
}

func handleScheduleButton(frontend Frontend, vendor models.Vendor) models.Vendor {
	switch frontend.button {
	case "Launch":
		vendor.IsDiscountOpen = true

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

	return vendor
}
