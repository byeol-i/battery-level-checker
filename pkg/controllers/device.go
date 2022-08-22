package controllers

import (
	"net/http"
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
// @Router /battery/{uuid} [get]
func (hdl *DeviceController) GetBattery(resp http.ResponseWriter, req *http.Request) {
	bytes := []byte(`
	{
		level: 20%,
		time: 
		status: charging,
	}`)
	resp.Write(bytes)

	resp.Header().Set("Content-Type", "text/html")
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

