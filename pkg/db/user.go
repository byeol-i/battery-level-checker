package db

import (
	"context"
	"errors"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/user"
)


func (db *Database) AddNewUser(userSpec user.UserImpl) error {
	const q = `
	INSERT INTO "User" ("id", "name") 
	VALUES ($1, $2)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := db.Conn.QueryRowContext(ctx, q, userSpec.Id, userSpec.Name)
	if res != nil {
		return errors.New(res.Err().Error())
	}
	return nil
}
