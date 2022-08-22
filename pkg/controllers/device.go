package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aglide100/battery-level-checker/pkg/models"
	"github.com/go-playground/validator"
)

type DeviceController struct {

}

func NewDeviceController() *DeviceController {
	return &DeviceController{}
}

// BatteryLevelChecker godoc
// @Summary Get Battery Level
// @Description Get devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param id path string true "Device ID"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=models.Device{}}
// @Router /battery/{deviceID} [get]
func (hdl *DeviceController) GetBattery(resp http.ResponseWriter, req *http.Request) {
	t := time.Now()
	
	mock := &models.Device{
		Name: req.URL.Path,
		Time: &t,
		BatteryLevel: 20,
		BatteryStatus: "charging",
	}

	respondJSON(resp, http.StatusOK, "success", mock)
}

// BatteryLevelChecker godoc
// @Summary Update Battery Level
// @Description Update devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param device body models.Device{} true "Device form"
// @Success 200 {object} models.JSONsuccessResult{}
// @Failure 400 {object} models.JSONfailResult{}
// @Router /battery/{deviceID} [put]
func (hdl *DeviceController) UpdateBattery(resp http.ResponseWriter, req *http.Request) {
	t, err := time.Parse("2006-01-02 15:04:05", req.PostFormValue("Time"))
	if err != nil {
		respondError(resp, 405, "time error")
	}

	bt, err := strconv.Atoi(req.PostFormValue("BatteryLevel"))
	if err != nil {
		respondError(resp, 405, "can't convert batteryLevel")
	}

	validDevice := models.Device{
		Name : req.PostFormValue("Name"),
		Time : &t,
		BatteryLevel: bt,
		BatteryStatus: req.PostFormValue("BatteryStatus"),
	}

	v := validator.New()

	err = v.Struct(validDevice)
	if err != nil {
		respondError(resp, 405, "not valid form")
	}

	respondJSON(resp, 200, "success", "")
} 

