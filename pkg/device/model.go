package device

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Device struct {
	DeviceInterface
	DeviceImpl
}

type DeviceInterface interface {
	GetDeviceSpec() *DeviceSpec
	GetProtoDeviceSpec() *pb_unit_device.Spec
	GetBatteryLevel() *BatteryLevel
	SetDeviceSpec(*DeviceSpec) error
	SetBatteryLevel(*BatteryLevel)
	SetDeviceId(string)
	GetDeviceId() string
	Validator() error
	Clone() *Device
	ToProtoDevice() (*pb_unit_device.Device)
	ToProtoBatteryLevel() (*pb_unit_device.BatteryLevel)
	ToJson() (string)
}

type DeviceImpl struct {
	Id string `validate:"required" json:"id" example:"f8aa8e64-eefa-423d-be40-a761231db093"`
	BatteryLevel BatteryLevel
	Spec DeviceSpec
}

type BatteryLevel struct {	
	Time BatteryTime `validate:"required" json:"time" example:"2006-01-02 15:04:05"`
	BatteryLevel  int        `validate:"required" json:"batteryLevel" example:"50"`
	BatteryStatus string     `validate:"required" json:"batteryStatus" example:"charging"`
}

type BatteryTime struct {
	time.Time
}
 
type DeviceSpec struct {
	Name       string `validate:"required" json:"name" example:"iphone 99xs"`
	Type       string `validate:"required" json:"type" example:"phone"`
	OS         string `validate:"required" json:"OS" example:"IOS"`
	OSversion  string `validate:"required" json:"OSversion" example:"99.192"`
	AppVersion string `validate:"required" json:"appVersion" example:"0.0.1"`
}

type Id struct {
	DeviceID   string `validate:"required" json:"deviceID" example:"f782fd74-d264-4901-8df5-c905f9df08db"`
}

func NewDevice() *Device {
	m := &Device{}

	return m
}

func Clone(d *Device) *Device {
	m := NewDevice()
	m.SetDeviceSpec(d.GetDeviceSpec())

	// m.SetDeviceSpec(d.GetDeviceSpec())
	return m

}

func (b *BatteryTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	logger.Info("time", zap.Any("t", s))

	t, err := dateparse.ParseAny(s)
	// layout := "2006-01-02 15:04:05"
	// 레이아웃을 여러개 두고 반복으로 돌리면서 파싱하는게 좋을듯
	// t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}

	b.Time = t
	
	return nil
}

func ParseFromJson(body string) (*BatteryLevel, error) {
	var newBatteryLevel BatteryLevel

	err := json.Unmarshal([]byte(body), &newBatteryLevel)
	if err != nil {
		return nil, err
	}

	return &newBatteryLevel, nil
}

func NewDeviceFromProto(pbDevice *pb_unit_device.Device) (*Device, error){
	spec := &DeviceSpec{
		Name: pbDevice.Spec.Name,
		Type: pbDevice.Spec.Type,
		OS: pbDevice.Spec.OS,
		OSversion: pbDevice.Spec.OsVersion,
		AppVersion: pbDevice.Spec.AppVersion,
	}

	err := SpecValidator(spec)
	if err != nil {
		return nil, err
	}
	
	m := NewDevice()
	m.SetDeviceSpec(spec)
	m.SetDeviceId(pbDevice.Id.Id)

	return m, nil
}

func (d *Device) SetDeviceId(id string) {
	d.Id=id
}

func (d *Device) GetDeviceId() string {
	return d.Id
}

func (d *Device) GetDeviceSpec() *DeviceSpec {
	return &d.Spec
}

func (d *Device) GetProtoDeviceSpec() *pb_unit_device.Spec {
	return &pb_unit_device.Spec{
		Name: d.Spec.Name,
		Type: d.Spec.Type,
		OS: d.Spec.OS,
		OsVersion: d.Spec.OSversion,
		AppVersion: d.Spec.AppVersion,
	}
}

func (d *Device) GetBatteryLevel() BatteryLevel {
	return d.BatteryLevel
}

func (d *Device) SetBatteryLevel(batteryLevel *BatteryLevel) {
	d.BatteryLevel = *batteryLevel
}

func (d *Device) SetDeviceSpec(spec *DeviceSpec) error {
	if spec == nil {
		return errors.New("spec is nil")
	}

	err := SpecValidator(spec)
	if err != nil {
		return err
	}
	d.Spec = *spec

	return nil
}

func (d *Device) ToJson() (string, error) {
	jsonData, err := json.Marshal(d.Spec)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (d *Device) ToProtoDevice() (*pb_unit_device.Device) {
	
	pbUnit := &pb_unit_device.Device{}

	id := &pb_unit_device.ID{
		Id: d.Id,
	}

	spec := &pb_unit_device.Spec{
		Name: d.Spec.Name,
		Type: d.Spec.Type,
		OS: d.Spec.OS,
		OsVersion: d.Spec.OSversion,
		AppVersion: d.Spec.AppVersion,
	}

	pbUnit.Id = id
	pbUnit.Spec = spec

	return pbUnit
}

func (d *Device) ToProtoBatteryLevel() (*pb_unit_device.BatteryLevel) {
	pbUnit := &pb_unit_device.BatteryLevel{}

	pbUnit.BatteryLevel = int64(d.BatteryLevel.BatteryLevel)
	pbUnit.BatteryStatus = d.BatteryLevel.BatteryStatus
	pbUnit.Time = timestamppb.New(d.BatteryLevel.Time.Time)

	return pbUnit
}

func ProtoToBatteryLevel(pbBatteryLevel *pb_unit_device.BatteryLevel) (*BatteryLevel, error) {
	const timeFormat = "2006-01-02 15:04:05"

	batteryLevel := &BatteryLevel{}
	batteryLevel.BatteryLevel = int(pbBatteryLevel.BatteryLevel)
	batteryLevel.BatteryStatus = pbBatteryLevel.BatteryStatus

	// parseTime, err := time.Parse(time.RFC3339, pbBatteryLevel.Time.AsTime())
    // if err != nil {
    //     return nil, err
    // }
	// batteryLevel.Time = &parseTime 


	if pbBatteryLevel.Time != nil {
		time := pbBatteryLevel.Time.AsTime()
		batteryLevel.Time = BatteryTime{Time:time}

	} else {
		logger.Error("battery time is null", zap.Any("pb bl", pbBatteryLevel))
	}

	return  batteryLevel, nil
}