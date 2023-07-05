package db

import (
	"context"
	"errors"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/google/uuid"
)

const maxRetry = 10;

func (db *Database) AddNewDevice(newDevice device.DeviceSpec, uid string) (string, error) {
	const selectQuery = `
	SELECT COUNT(*) FROM "Device"
	WHERE "name" = $1 AND "type" = $2 AND "os_name" = $3 AND "os_version" = $4 AND "app_version" = $5 AND "uid" = $6;	
	`
	
	const insertQuery = `
	INSERT INTO "Device" 
	("device_id", "name", "type", "os_name", "os_version", "app_version", "uid") 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	
	err := device.SpecValidator(&newDevice)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	tx, err := db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(ctx, selectQuery, newDevice.Name, newDevice.Type, newDevice.OS, newDevice.OSversion, newDevice.AppVersion, uid).Scan(&count)
	if err != nil {
		return "", err
	}

	if count > 0 {
		return "",errors.New("device already exists")
	}
	
	deviceID := ""

	for i := 0; i < maxRetry; i++ {
		newUUID := uuid.New().String()

		var existingCount int
		err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM \"Device\" WHERE \"device_id\" = $1", newUUID).Scan(&existingCount)
		if err != nil {
			return "", err
		}

		if existingCount == 0 {
			deviceID = newUUID
			break
		}
	}

	if deviceID == "" {
		return "", errors.New("failed to generate unique device ID")
	}

	_, err = tx.ExecContext(ctx, insertQuery, deviceID, newDevice.Name, newDevice.Type, newDevice.OS, newDevice.OSversion, newDevice.AppVersion, uid)
	if err != nil {
		return "", err
	}

	return deviceID, tx.Commit()
}

func (db *Database) RemoveDevice(deviceId device.Id, uid string) error {
	const q = `
	DELETE FROM "Device" 
	WHERE "device_id" = $1 AND
	"uid" = $2
	`

	ctx := context.Background()
	tx, err := db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.QueryContext(ctx, q, deviceId.DeviceID, uid)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db *Database) GetDevices(uid string) ([]*device.Device, error) {
	const q = `
	SELECT ("device_id", "name", "type", "os_name", "os_version", app_version) 
	FROM "Device" 
	WHERE "uid" = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var devices []*device.Device
	
	rows, err := db.Conn.QueryContext(ctx, q, uid)
	if err != nil {
		return nil, err
	}

	d := device.NewDevice()

	for rows.Next() {
		var (
			DeviceId string
			Name       string 
			Type       string 
			OS         string 
			OSversion  string 
			AppVersion string 
		)

		err := rows.Scan(
			&DeviceId,
			&Name,
			&Type,
			&OS,
			&OSversion,
			&AppVersion,
		)

		if err != nil {
			return nil, err
		}

		
		d.SetDeviceSpec(
			&device.DeviceSpec{
				Name: Name,
				Type: Type,
				OS: OS,
				OSversion: OSversion,
				AppVersion: AppVersion,
			},
		)

		d.SetDeviceId(DeviceId)

		devices = append(devices, d.Clone())
	}

	return devices, nil
}

func (db *Database) GetDevice(deviceId string, uid string) (*device.DeviceSpec, error) {
	const q = `
	SELECT * FROM "Device"
	WHERE "uid" = $1 AND
	"device_id" = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := db.Conn.QueryRowContext(ctx, q, uid, deviceId)
	if res != nil {
		return nil, errors.New(res.Err().Error())
	}

	newDevice := device.NewDevice()

	var (
		Name       string 
		Type       string 
		OS         string 
		OSversion  string 
		AppVersion string 
	)

	err := res.Scan(
		&Name,
		&Type,
		&OS,
		&OSversion,
		&AppVersion,
	)

	if err != nil {
		return nil, err
	}

	newDevice.SetDeviceSpec(
		&device.DeviceSpec{
			Name: Name,
			Type: Type,
			OS: OS,
			OSversion: OSversion,
			AppVersion: AppVersion,
		},
	)

	err = device.SpecValidator(newDevice.GetDeviceSpec())
	if err != nil {
		return nil, err
	}

	return newDevice.GetDeviceSpec(), nil
}
