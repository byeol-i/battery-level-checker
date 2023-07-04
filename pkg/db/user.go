package db

import (
	"context"
	"strings"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/user"
)

func (db *Database) AddNewUser(userSpec user.UserImpl, userCredential user.UserCredential) error {
	const q = `
	INSERT INTO "User" ("uid", "name", "email") 
	VALUES ($1, $2, $3)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res := db.Conn.QueryRowContext(ctx, q, strings.Replace(userCredential.Uid, "\"", "", -1), userSpec.Name, userSpec.Email)
	if res != nil {
		return res.Err()
	}

	return nil
}
