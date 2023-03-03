package device

import (
	"time"

	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
)

type Device struct {
	DeviceInterface
}

type DeviceImpl struct {
	Id string
	BatteryLevel BatteryLevel
	Spec Spec
}

type DeviceInterface interface {
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
	ToProtoModel
	// PushToken  string `json:"pushToken" example:"abcd"`
}

type ToProtoModel interface {
	ProtoSpec() pb_unit_device.Spec
	ProtoBatteryLevel(BatteryLevel) pb_unit_device.BatteryLevel
}

func (spec Spec) ProtoSpec() pb_unit_device.Spec {
	return pb_unit_device.Spec{
		Name: spec.Name,
		Type: spec.Type,
		Os: spec.OS,
		OsVersion: spec.OSversion,
		AppVersion: spec.AppVersion,
	}
}

type Id struct {
	DeviceID   string `validate:"required" json:"deviceID" example:"f782fd74-d264-4901-8df5-c905f9df08db"`
}

func NewDevice() *Device {
	m := &Device{}

	return m
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

