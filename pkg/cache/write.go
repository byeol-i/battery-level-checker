package cache

import "github.com/patrickmn/go-cache"

func (m *CacheManager) Write(deviceId, userId string, value []byte) {
	m.deviceCache.Delete(deviceId)
	m.deviceCache.Set(deviceId, value, cache.DefaultExpiration)

	res, found := m.userCache.Get(userId)
	if found {
		list := AddDeviceInUser(res.([]string), m.deviceCache, deviceId)
		 
		m.userCache.Replace(userId, list, cache.DefaultExpiration)
	} else {
		newList := []string{deviceId}

		m.userCache.Set(userId, newList, cache.DefaultExpiration)
	}
}

func AddDeviceInUser(list []string, deviceCache *cache.Cache, newDeviceId string) []string {
	var newList []string
	
	find := false
	for _, deviceId := range list {
		_, found := deviceCache.Get(deviceId)

		if (found) {
			newList = append(newList, deviceId)
		}
		
		if deviceId == newDeviceId {
			find = true
		}
	}

	if !find {
		newList = append(newList, newDeviceId)
	}

	return newList
}
