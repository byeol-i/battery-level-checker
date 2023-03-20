package db

import (
	"fmt"
	"log"

	"github.com/byeol-i/battery-level-checker/pkg/device"
)

func (db *Database) GetBattery(deviceId string, uid string) (*device.BatteryLevel, error) {
	const q = `
	SELECT * FROM "BatteryLevel"
	WHERE "device_id" = $1 AND
	"user_id" = $2	
	ORDER BY time DESC
	LIMIT 1;
	`
	batteryLevel := &device.BatteryLevel{}
	err := db.Conn.QueryRow(q, deviceId, uid).Scan(&batteryLevel.Time, &batteryLevel.BatteryLevel, &batteryLevel.BatteryStatus)
	if err != nil {
		return nil, err
	}

	return batteryLevel, nil
}

func (db *Database) GetUsersAllBatteryLevels(uid string) ([]*device.BatteryLevel, error) {
	var batteryLevels []*device.BatteryLevel

	q := `
	SELECT * FROM "BatteryLevel"
	WHERE "user_id" = $1`

	rows, err := db.Conn.Query(q, uid)
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

func (db *Database) GetAllBatteryLevels(deviceId string, uid string) ([]*device.BatteryLevel, error) {
	var batteryLevels []*device.BatteryLevel

	q := `
	SELECT * FROM "BatteryLevel"
	WHERE "device_id" = $1 AND
	"user_id" = $2`

	rows, err := db.Conn.Query(q, deviceId, uid)
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

func (db *Database) UpdateBattery(deviceId string, uid string, batteryLevel *device.BatteryLevel) error {
	const q = `
	INSERT INTO "BatteryLevel"("device_id", "user_id", "time", "battery_level", "battery_status")
	VALUES ($1, $2, $3, $4, $5)
	`

	err := device.BatteryLevelValidator(batteryLevel)
	if err != nil {
		return err
	}

	_, err = db.Conn.Exec(q, deviceId, uid, batteryLevel.Time, batteryLevel.BatteryLevel, batteryLevel.BatteryStatus)
	if err != nil {
		return err
	}

	return nil
}