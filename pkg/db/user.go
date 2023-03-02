package db

import (
	"context"
	"errors"
	"time"
)


func (db *Database) AddNewUser(uid string, name string) error {
	const q = `
	INSERT INTO "User" ("id", "name") 
	VALUES ($1, $2)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := db.Conn.QueryRowContext(ctx, q, uid, name)
	if res != nil {
		return errors.New(res.Err().Error())
	}
	return nil
}
