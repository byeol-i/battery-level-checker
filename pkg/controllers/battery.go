package controllers

import (
	"net/http"
	"regexp"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	// "github.com/byeol-i/battery-level-checker/pkg/models"
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
	pattern := regexp.MustCompile(`/device/(\w+)`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
        http.NotFound(resp, req)
        return
    }

	uid := req.Header.Get("Uid")

	res, err := dbSvc.CallGetBattery(matches[1], uid)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
	}

	// consumer.GetTopics()
	respondJSON(resp, http.StatusOK, "success", res)
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
func (hdl *BatteryController) GetAllBattery(resp http.ResponseWriter, req *http.Request) {
	
	pattern := regexp.MustCompile(`/battery/(\w+)`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
        http.NotFound(resp, req)
        return
    }

	uid := req.Header.Get("Uid")

	res, err := dbSvc.CallGetAllBattery(matches[1], uid)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
	}

	// consumer.GetTopics()
	respondJSON(resp, http.StatusOK, "success", res)
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

	// producer.Write() 

	respondJSON(resp, 200, "success", "")
}
