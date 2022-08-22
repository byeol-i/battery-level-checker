package controllers

import (
	"net/http"
	"time"

	"github.com/aglide100/battery-level-checker/pkg/models"
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
// @Failure 400 
// @Success 200 {object} models.Device{}
// @Router /battery/{uuid} [get]
func (hdl *DeviceController) GetBattery(resp http.ResponseWriter, req *http.Request) {
	t := time.Now()
	
	mock := &models.Device{
		Name: "",
		Time: &t,
		BatteryLevel: 20,
		BatteryStatus: "charging",
	}

	respondJSON(resp, http.StatusOK, mock)
}

// BatteryLevelChecker godoc
// @Summary Update Battery Level
// @Description Update devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Failure 400 
// @Router /battery/{uuid} [put]
func (hdl *DeviceController) UpdateBattery(resp http.ResponseWriter, req *http.Request) {

} 

