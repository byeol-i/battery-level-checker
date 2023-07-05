package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	"go.uber.org/zap"
)

type DeviceControllers struct {
	basePattern string
}

func NewDeviceController(basePattern string) *DeviceControllers {
	return &DeviceControllers{
		basePattern: basePattern,
	}
}

// AddDevice godoc
// @Summary Register New Device
// @Description Register New Device
// @Tags Device
// @Accept json
// @Produce json
// @Param deviceInfo body device.DeviceImpl true "add device"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /device [post]
func (hdl *DeviceControllers) AddNewDevice(resp http.ResponseWriter, req *http.Request) {
	var deviceSpec device.DeviceSpec

	err := json.NewDecoder(req.Body).Decode(&deviceSpec)
	if err != nil {
		logger.Error("Json parse error", zap.Error(err))
		respondError(resp, http.StatusRequestedRangeNotSatisfiable, "invalid format")
		return
	}

	err = device.SpecValidator(&deviceSpec)
	if err != nil {
		logger.Error("Device spec validator error", zap.Error(err))
		respondError(resp, http.StatusRequestedRangeNotSatisfiable, "invalid format")
		return
	}

	newDevice := device.NewDevice()
	err = newDevice.SetDeviceSpec(&deviceSpec)
	if err != nil {
		logger.Error("Device spec validator error", zap.Error(err))
		respondError(resp, http.StatusRequestedRangeNotSatisfiable, "invalid format")
		return
	}

	uid := req.Header.Get("Uid")
	logger.Info("Add New device", zap.String("uid", uid))
	
	err = dbSvc.CallAddNewDevice(newDevice, strings.Replace(uid, "\"", "", -1))
	if err != nil {
		logger.Error("dbSvc's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
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
	pattern := regexp.MustCompile(hdl.basePattern + `/device/([a-zA-Z0-9-]+)`)
	
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches[1]) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }

	uid := req.Header.Get("Uid")
	
	err := dbSvc.CallRemoveDevice(matches[1], uid)
	if err != nil {
		logger.Error("dbSvc's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}
	
	respondJSON(resp, http.StatusOK, "success", nil)
}

// GetDevices godoc
// @Summary Get Devices
// @Description Get Devices
// @Tags Device
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=device.DeviceImpl{}}
// @Router /device/ [get]
func (hdl *DeviceControllers) GetDevices(resp http.ResponseWriter, req *http.Request) {
	uid := req.Header.Get("Uid")
	
	res, err := dbSvc.CallGetAllDevices(uid)
	if err != nil {
		logger.Error("dbSvc's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}
	
	respondJSON(resp, http.StatusOK, "", res)
}