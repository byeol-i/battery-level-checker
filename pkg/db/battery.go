package db

import (
	"fmt"
	"log"

	"github.com/byeol-i/battery-level-checker/pkg/device"
)

func (db *Database) GetBattery(deviceId string) (*device.BatteryLevel, error) {
	const q = `
	SELECT * FROM "BatteryLevel"
	WHERE "deviceId" = $1;	
	ORDER BY time DESC
	LIMIT 1;
	`
	batteryLevel := &device.BatteryLevel{}
	err := db.Conn.QueryRow(q).Scan(&batteryLevel.Time, &batteryLevel.BatteryLevel, &batteryLevel.BatteryStatus)
	if err != nil {
		return nil, err
	}

	return batteryLevel, nil
}

func (db *Database) GetAllBatteryLevels(deviceId string) ([]*device.BatteryLevel, error) {
	var batteryLevels []*device.BatteryLevel
	q := `
	SELECT * FROM "BatteryLevel"
	WHERE "deviceId" = $1`
	rows, err := db.Conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		batteryLevel := &device.BatteryLevel{}
		err := rows.Scan(&batteryLevel.Time, &batteryLevel.BatteryLevel, &batteryLevel.BatteryStatus)
		if err != nil {
			log.Printf("failed to scan row: %v", err)
			continue
		}
		batteryLevels = append(batteryLevels, batteryLevel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate rows: %v", err) 
	}

	return batteryLevels, nil
}
