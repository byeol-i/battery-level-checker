package db

import (
	"context"
	"errors"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/models"
)

func (db *Database) AddNewDevice(device models.DeviceDetail) error {
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
