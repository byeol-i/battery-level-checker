package controllers

import (
	"net/http"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/models"
	"github.com/byeol-i/battery-level-checker/pkg/producer"
)

type BatteryController struct {

}

func NewBatteryController() *BatteryController {
	return &BatteryController{}
}

// GetBatteryLevel godoc
// @Summary Get Battery Level
// @Description Get devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param id path string true "Device ID"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=models.Device{}}
// @Router /battery/{deviceID} [get]
func (hdl *BatteryController) GetBattery(resp http.ResponseWriter, req *http.Request) {
	t := time.Now()
	
	mock := &models.Device{
		Name: req.URL.Path,
		Time: &t,
		BatteryLevel: 20,
		BatteryStatus: "charging",
	}
	
	// consumer.GetTopics()
	respondJSON(resp, http.StatusOK, "success", mock)
}

// GetBatteryList godoc
// @Summary Get All Device's Battery Level
// @Description Get devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{} 
// @Success 200 {object} models.JSONsuccessResult{data=[]models.Device{}}
// @Router /battery/ [get]
func (hdl *BatteryController) GetBatteryList(resp http.ResponseWriter, req *http.Request) {
	t := time.Now()
	
	mock1 := &models.Device{
		Name: "test1",
		Time: &t,
		BatteryLevel: 20,
		BatteryStatus: "charging",
	}

	mock2 := &models.Device{
		Name: "test2",
		Time: &t,
		BatteryLevel: 20,
		BatteryStatus: "charging",
	}

	mock3 := &models.Device{
		Name: "test3",
		Time: &t,
		BatteryLevel: 20,
		BatteryStatus: "charging",
	}
	
	data := []*models.Device{
		mock1, 
		mock2,
		mock3,
	}
	
	// consumer.GetTopics()
	respondJSON(resp, http.StatusOK, "success", data)
}

// UpdateBatteryLevel godoc
// @Summary Update Battery Level
// @Description Update devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param device body models.Device true "Battery input form"
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} models.JSONsuccessResult{}
// @Failure 400 {object} models.JSONfailResult{}
// @Router /battery/{deviceID} [put]
func (hdl *BatteryController) UpdateBattery(resp http.ResponseWriter, req *http.Request) {
	// t, err := time.Parse("2006-01-02 15:04:05", req.PostFormValue("Time"))
	// if err != nil {
	// 	respondError(resp, 405, "time error")
	// }

	// bt, err := strconv.Atoi(req.PostFormValue("BatteryLevel"))
	// if err != nil {
	// 	respondError(resp, 405, "can't convert batteryLevel")
	// }

	// validDevice := models.Device{
	// 	Name : req.PostFormValue("Name"),
	// 	Time : &t,
	// 	BatteryLevel: bt,
	// 	BatteryStatus: req.PostFormValue("BatteryStatus"),
	// }

	// v := validator.New()

	// err = v.Struct(validDevice)
	// if err != nil {
	// 	respondError(resp, 405, "not valid form")
	// }
	logger.Info("Update Battery")

	producer.Write()

	respondJSON(resp, 200, "success", "")
} 