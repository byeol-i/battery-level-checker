package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/byeol-i/battery-level-checker/pkg/device"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
)

type DeviceControllers struct {

}

func NewDeviceController() *DeviceControllers {
	return &DeviceControllers{}
}

// AddDevice godoc
// @Summary Register New Device
// @Description Register New Device
// @Tags Device
// @Accept json
// @Produce json
// @Param deviceInfo body models.Spec true "add device"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /device [post]
func (hdl *DeviceControllers) AddNewDevice(resp http.ResponseWriter, req *http.Request) {
	respondJSON(resp, http.StatusOK, "success", nil)
}

// DeleteDevice godoc
// @Summary Delete Device
// @Description Delete Device
// @Tags Device
// @Accept json
// @Produce json
// @Param id path string true "Device ID"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /device/{deviceID} [delete]
func (hdl *DeviceControllers) DeleteDevice(resp http.ResponseWriter, req *http.Request) {
	
	var deviceSpec device.DeviceSpec

	err := json.NewDecoder(req.Body).Decode(&deviceSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, "invalid format")
	}

	err = device.SpecValidator(&deviceSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())	
	}

	newDevice := device.NewDevice()
	newDevice.SetDeviceSpec(&deviceSpec)
	
	err = dbSvc.CallAddNewDevice(newDevice)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
	}

	respondJSON(resp, http.StatusOK, "success", nil)
}