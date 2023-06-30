package controllers

import (
	"net/http"
	"regexp"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	"go.uber.org/zap"
	// "github.com/byeol-i/battery-level-checker/pkg/models"
)

type BatteryController struct {
	basePattern string
}

func NewBatteryController(basePattern string) *BatteryController {
	return &BatteryController{
		basePattern: basePattern,
	}
}

// GetUsersAllBattery godoc
// @Summary Get User's all Battery Level
// @Description Get devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=device.DeviceImpl{}}
// @Router /battery/ [get]
func (hdl *BatteryController) GetUsersAllBattery(resp http.ResponseWriter, req *http.Request) {
	uid := req.Header.Get("Uid")

	
	res, err := dbSvc.CallGetUsersAllBattery(uid)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())
		return
	}

	// consumer.GetTopics()
	respondJSON(resp, http.StatusOK, "success", res)
}

// GetHistoryAllBattery godoc
// @Summary Get All Device's Battery Level
// @Description Get devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=[]device.DeviceImpl{}}
// @Router /battery/history/{deviceId}} [get]
func (hdl *BatteryController) GetHistoryAllBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(`/battery/history/(\w+)`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }

	uid := req.Header.Get("Uid")
	logger.Info("GetHistoryAllBattery", zap.String("uid", uid), zap.String("device", matches[1]))
	
	res, err := dbSvc.CallGetAllBattery(matches[1], uid)
	if err != nil {
		logger.Error("dbSvc's error", zap.Error(err))
		respondError(resp, http.StatusBadRequest, err.Error())
		return
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
// @Param device body device.DeviceImpl true "Battery input form"
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} models.JSONsuccessResult{}
// @Failure 400 {object} models.JSONfailResult{}
// @Router /battery/{deviceID} [put]
func (hdl *BatteryController) UpdateBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(`/battery/history/(\w+)`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }

	// t, err := time.Parse("2006-01-02 15:04:05", req.PostFormValue("Time"))
	// if err != nil {
	// 	respondError(resp, 405, "time error")
	// }

	// bt, err := strconv.Atoi(req.PostFormValue("BatteryLevel"))
	// if err != nil {
	// 	respondError(resp, 405, "can't convert batteryLevel")
	// }

	// validDevice := device.DeviceImpl{
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
