package db

import (
	"context"
	"errors"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/device"
)

func (db *Database) AddNewDevice(newDevice device.DeviceSpec, uid string) error {
	const selectQuery = `
	SELECT COUNT(*) FROM "Device"
	WHERE "name" = $1 AND "type" = $2 AND "os_name" = $3 AND "os_version" = $4 AND "app_version" = $5 AND "user_id" = $6;	
	`
	
	const insertQuery = `
	INSERT INTO "Device" 
	("name", "type", "os_name", "os_version", "app_version", "user_id") 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING "id";
	`
	
	err := device.SpecValidator(&newDevice)
	if err != nil {
		return err
	}

	ctx := context.Background()
	tx, err := db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(ctx, selectQuery, newDevice.Name, newDevice.Type, newDevice.OS, newDevice.OSversion, newDevice.AppVersion, uid).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("device already exists")
	}

	var returnedId = device.Id{}

	err = tx.QueryRowContext(ctx, insertQuery, newDevice.Name, newDevice.Type, newDevice.OS, newDevice.OSversion, newDevice.AppVersion, uid).Scan(&returnedId)
	if err != nil {
		return err
	}

	// admin, err := consumer.GetAdmin()
	// if err != nil {
	// 	return err
	// }

	// err = consumer.CreateTopic(admin, uid)
	// if err != nil {
	// 	return err
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	// res := db.Conn.QueryRowContext(ctx, q, newDevice.Name, newDevice.Type, newDevice.OS, newDevice.OSversion, newDevice.AppVersion, uid)
	// if res != nil {
	// 	return errors.New(res.Err().Error())
	// }
	return tx.Commit()
}


func (db *Database) RemoveDevice(deviceId device.Id, uid string) error {
	const q = `
	DELETE FROM "Device" 
	WHERE "device_id" = $1 AND
	"user_id" = $2
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
	SELECT * FROM "Device" 
	WHERE "user_id" = $1
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
			Name       string 
			Type       string 
			OS         string 
			OSversion  string 
			AppVersion string 
		)

		err := rows.Scan(
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

		devices = append(devices, d.Clone())
	}

	return devices, nil
}

func (db *Database) GetDevice(deviceId string, uid string) (*device.DeviceSpec, error) {
	const q = `
	SELECT * FROM "Device"
	WHERE "user_id" = $1 AND
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
