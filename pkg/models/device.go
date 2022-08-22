package models

import "time"

type Device struct {
	Name string `json:"Name"`
	Time *time.Time `json:"Time"`
	BatteryLevel int `json:"BatteryLevel"`
	BatteryStatus string `json:"BatteryStatus"`
}