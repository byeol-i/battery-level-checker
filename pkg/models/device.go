package models

import (
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/logger"

	// pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

type Device struct {
	DeviceInfo
	// DeviceImpl
	// Spec
	// BatteryLevel
	// Spec Spec
	// BatteryLevel BatteryLevel
	// Validator error
}

type DeviceImpl struct {
	Id string
	BatteryLevel BatteryLevel
	Spec Spec
}

type DeviceInfo interface {
	GetDeviceSpec() Spec
	GetBatteryLevel() BatteryLevel
	SetSpec(Spec)
	SetBatteryLevel(BatteryLevel)
	Validator() error
}

type BatteryLevel struct {	
	Time          *time.Time `validate:"required" json:"time" example:"2006-01-02 15:04:05"`
	BatteryLevel  int        `validate:"required" json:"batteryLevel" example:"50"`
	BatteryStatus string     `validate:"required" json:"batteryStatus" example:"charging"`
}

type Spec struct {
	Name       string `validate:"required" json:"name" example:"iphone 99xs"`
	Type       string `validate:"required" json:"type" example:"phone"`
	OS         string `validate:"required" json:"OS" example:"IOS"`
	OSversion  string `validate:"required" json:"OSversion" example:"99.192"`
	AppVersion string `validate:"required" json:"appVersion" example:"0.0.1"`
	// PushToken  string `json:"pushToken" example:"abcd"`
}

type Id struct {
	DeviceID   string `validate:"required" json:"deviceID" example:"f782fd74-d264-4901-8df5-c905f9df08db"`
}

func NewDevice() *Device {
	m := &Device{}

	return m
}

func SpecValidator(spec *Spec) error {
	var validate *validator.Validate

	err := validate.Struct(spec)
	if err != nil {
		logger.Error("Device's Spec is not valid", zap.Any("spec", spec))
		return err
	}

	return nil
}


func BatteryLevelValidator(level *BatteryLevel) error {
	var validate *validator.Validate

	err := validate.Struct(level)
	if err != nil {
		logger.Error("Device's BatteryLevel is not valid", zap.Any("BatteryLevel", level))
		return err
	}

	return nil
}


func (d *DeviceImpl) GetSpec() Spec {
	return d.Spec
}

func (d *DeviceImpl) GetBatteryLevel() BatteryLevel {
	return d.BatteryLevel
}


func (d *DeviceImpl) SetBatteryLevel(batteryLevel BatteryLevel) {
	d.BatteryLevel = batteryLevel
}

func (d *DeviceImpl) SetSpec(spec Spec) {
	d.Spec = spec
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
