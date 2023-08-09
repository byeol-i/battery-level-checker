package cache

import (
	"encoding/json"

	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)
type BatteryLevelWrapper struct {
	device.Id
	device.BatteryLevel
}

func (m *CacheManager) GetDataFromUserId(userId string) []string {
	deviceList, found := m.userCache.Get(userId) 
	if !found {
		return nil
	}

	return GetResultFromDeviceList(deviceList.([]string), m.deviceCache)
}

func GetResultFromDeviceList(deviceList []string, deviceCache *cache.Cache) []string {
	var result []string
	for _, deviceId := range deviceList {
		res, found := deviceCache.Get(deviceId)

		if found {
			bl := &device.BatteryLevel{}
			err := json.Unmarshal(res.([]byte), bl)
			if err != nil {
				logger.Error("Can't Unmarshal", zap.Error(err))
				continue
			}
			
			blWrapper := &BatteryLevelWrapper{
				device.Id{DeviceID: deviceId},
				*bl,
			}

			marshaledData, err := json.Marshal(blWrapper)
			if err != nil {
				logger.Error("Can't Marshal data", zap.Error(err))
				continue
			}

			result = append(result, string(marshaledData[:]))
		}
	}

	return result
}