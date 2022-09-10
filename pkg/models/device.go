package models

import (
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

type Device struct {
	Name string `validate:"required" json:"name" example:"iphone 99xs"` 
	Time *time.Time `validate:"required" json:"time" example:"2006-01-02 15:04:05"`
	BatteryLevel int `validate:"required" json:"batteryLevel" example:"50"`
	BatteryStatus string `validate:"required" json:"batteryStatus" example:"charging"`
}

type DeviceDetail struct {
	Name string `json:"name" example:"iphone 99xs"`
	Type string `json:"type" example:"phone"`
	OS string `json:"OS" example:"IOS"`
	OSversion string `json:"OSversion" example:"99.192"`
	AppVersion string `json:"appVersion" example:"0.0.1"`
	PushToken string `json:"pushToken" example:"abcd"`
}

func DeviceValidation(device *Device) error {
	var validate *validator.Validate

	err := validate.Struct(device)
	if err != nil {
		logger.Error("Device models is not valid", zap.Any("device", device))
		return err
	}

	return nil
}

// func DeviceDetailValidation(deviceDetail *DeviceDetail) error {
// 	var validate *validator.Validate

// 	err := validate.Struct(deviceDetail)
// 	if err != nil {
// 		logger.Error("deviceDetail models is not valid", zap.Any("deviceDetail", deviceDetail))
// 		return err
// 	}

// 	return nil
// }