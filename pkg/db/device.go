package db

import (
	"context"
	"errors"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/models"
)

func (db *Database) AddNewDevice(device models.Device) error {
	const q = `
	INSERT INTO Device (
		"Id", "Name", "",
	)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := db.Conn.QueryRowContext(ctx, q)
	if res != nil {
		return errors.New(res.Err().Error())
	}
	return nil
}


func (db *Database) RemoveDevice(device models.Id) error {
	const q = `
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := db.Conn.QueryRowContext(ctx, q)
	if res != nil {
		return errors.New(res.Err().Error())
	}
	return nil
}

func (db *Database) GetDevices(uid string) (*[]models.Device, error) {
	const q = `
	SELECT * FROM Device 
	WHERE "userId" = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	rows, err := db.Conn.QueryContext(ctx, q, uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

	}

	return nil, nil
}

func (db *Database) GetDevice(deviceId string) (*models.Device, error) {
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
