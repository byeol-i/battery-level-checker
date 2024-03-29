package controllers

import (
	"io"
	"net/http"
	"regexp"

	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/producer"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"

	cacheSvc "github.com/byeol-i/battery-level-checker/pkg/svc/cache"
	"go.uber.org/zap"
	// "github.com/byeol-i/battery-level-checker/pkg/models"
)

type BatteryController struct {
	basePattern string
	dbClient *dbSvc.DBSvcClient
	cacheClient *cacheSvc.CacheSvcClient
}

func NewBatteryController(basePattern string, dbClientAddr, cacheSvcAddr string,) *BatteryController {
	return &BatteryController{
		basePattern: basePattern,
		dbClient: dbSvc.NewDBSvcClient(dbClientAddr),
		cacheClient: cacheSvc.NewCacheSvcClient(cacheSvcAddr),
	}
}

// GetUsersCachedBattery godoc
// @Summary Get User's all Battery Level
// @Description Get device's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=[]device.DeviceImpl{Id,BatteryLevel}}
// @Router /battery/ [get]
func (hdl *BatteryController) GetUsersCachedBattery(resp http.ResponseWriter, req *http.Request) {
	uid := req.Header.Get("Uid")

	res, err := hdl.cacheClient.CallGetCurrentMsg(uid)
	
	// res, err := hdl.dbClient.CallGetUsersAllBattery(uid)
	if err != nil {
		logger.Error("dbClient's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondJSON(resp, http.StatusOK, "GetUsersAllBattery", res)
}


// GetSpecificDevice's battery godoc
// @Summary Get Specific Device's Battery Level
// @Description Get device battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param id path string true "Device ID"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=device.DeviceImpl{Id,BatteryLevel}}
// @Router /battery/ [get]
func (hdl *BatteryController) GetBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(hdl.basePattern + `/battery/([a-zA-Z0-9-]+)`)
	
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) == 0 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }
	
	uid := req.Header.Get("Uid")

	res, err := hdl.dbClient.CallGetBattery(matches[1], uid)
	if err != nil  {
		logger.Error("dbClient's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}

	// err := consumer.ConsumeLatestMessage("battery_device_"+uid+"_"+matches[1])
	// if err != nil {
	// 	logger.Error("Can't consume msg", zap.Error(err))
	// }

	// respondJSON(resp, http.StatusOK, "GetBattery", res)

	respondJSON(resp, http.StatusOK, "GetBattery", res)
}



// GetHistoryAllBattery godoc
// @Summary Get All Device's Battery Level
// @Description Get devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=[]device.DeviceImpl{Id,BatteryLevel}}
// @Router /battery/history/{deviceId}} [get]
func (hdl *BatteryController) GetHistoryAllBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(`/battery/history/(\w+)`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches[1]) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }

	uid := req.Header.Get("Uid")
	res, err := hdl.dbClient.CallGetAllBattery(matches[1], uid)
	if err != nil {
		logger.Error("dbClient's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondJSON(resp, http.StatusOK, "GetHistoryAllBattery", res)
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
// @Router /battery/{deviceID} [post]
func (hdl *BatteryController) UpdateBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(hdl.basePattern + `/battery/([a-zA-Z0-9-]+)`)
	
    deviceId := pattern.FindStringSubmatch(req.URL.Path)
	if len(deviceId) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }
	
	body, err := io.ReadAll(req.Body)
	if err != nil {
		logger.Error("Can't reading body", zap.Error(err))
	}

	newBatteryLevel, err := device.ParseFromJson(string(body[:]))
	if err != nil {
		logger.Error("Json parse error", zap.Error(err))
		respondError(resp, http.StatusRequestedRangeNotSatisfiable, "invalid format")
		return
	}

	err = device.BatteryLevelValidator(newBatteryLevel)
	if err != nil {
		logger.Error("Battery level is not validate", zap.Error(err))	
		respondError(resp, http.StatusRequestedRangeNotSatisfiable, "invalid format")
		return
	}


	uid := req.Header.Get("Uid")

	// err = hdl.dbClient.CallUpdateBatteryLevel(deviceId[1], uid, newBatteryLevel)
	// if err != nil {
	// 	logger.Error("dbClient's error", zap.Error(err))
	// 	respondError(resp, http.StatusInternalServerError, "Internal server error")
	// 	return
	// }

	err = producer.WriteBatteryTime(newBatteryLevel, deviceId[1], uid)
	if err != nil {
		logger.Error("can't producer ", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondJSON(resp, 200, "UpdateBattery", "")
}
