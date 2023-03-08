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
	WHERE "name" = $1 AND "type" = $2 AND "osName" = $3 AND "osVersion" = $4 AND "appVersion" = $5 AND "userId" = $6;	
	`
	
	const insertQuery = `
	INSERT INTO "Device" 
	("name", "type", "osName", "osVersion", "appVersion", "userId") 
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
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	// res := db.Conn.QueryRowContext(ctx, q, newDevice.Name, newDevice.Type, newDevice.OS, newDevice.OSversion, newDevice.AppVersion, uid)
	// if res != nil {
	// 	return errors.New(res.Err().Error())
	// }
	return tx.Commit()
}


func (db *Database) RemoveDevice(deviceId device.Id) error {
	const q = `
	DELETE FROM "Device" 
	WHERE "id" = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res, err := db.Conn.ExecContext(ctx, q, deviceId)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected while deleting device with ID")
	}

	return nil
}

func (db *Database) GetDevices(uid string) ([]*device.Device, error) {
	const q = `
	SELECT * FROM "Device" 
	WHERE "userId" = $1
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

func (db *Database) GetDevice(deviceId string) (*device.DeviceSpec, error) {
	const q = `
	SELECT * FROM "Device"
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := db.Conn.QueryRowContext(ctx, q, deviceId)
	if res != nil {
		return nil, errors.New(res.Err().Error())
	}
	return nil, nil
}
