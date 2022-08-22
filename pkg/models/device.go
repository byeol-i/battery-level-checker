package models

import "time"

type Device struct {
	Name string `validate:"required" json:"name" example:"iphone 99xs"` 
	Time *time.Time `validate:"required" json:"time" example:"2006-01-02 15:04:05"`
	BatteryLevel int `validate:"required" json:"batteryLevel" example:"50"`
	BatteryStatus string `validate:"required" json:"batteryStatus" example:"charging"`
}
