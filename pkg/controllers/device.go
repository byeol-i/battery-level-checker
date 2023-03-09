package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	"go.uber.org/zap"
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
	// deviceID := req.URL.Path
	// var matches []string
	pattern := regexp.MustCompile(`/device/(\w+)`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
        http.NotFound(resp, req)
        return
    }

	logger.Info("matches", zap.Any("1", matches))
	err := dbSvc.CallRemoveDevice(matches[1])
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
		return
	}
	
	respondJSON(resp, http.StatusOK, "success", nil)
}