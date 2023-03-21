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
		return
	}

	err = device.SpecValidator(&deviceSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())	
		return
	}

	newDevice := device.NewDevice()
	err = newDevice.SetDeviceSpec(&deviceSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
		return
	}

	uid := req.Header.Get("Uid")
	logger.Info("Add New device", zap.String("uid", uid))
	

	err = dbSvc.CallAddNewDevice(newDevice, uid)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
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
	pattern := regexp.MustCompile(hdl.basePattern + `/device/(\\w+)`)
	
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }

	uid := req.Header.Get("Uid")
	logger.Info("Delete device", zap.String("uid", uid), zap.String("device", matches[1]))
	
	err := dbSvc.CallRemoveDevice(matches[1], uid)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
		return
	}
	
	respondJSON(resp, http.StatusOK, "success", nil)
}