package db

import (
	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
	"github.com/byeol-i/battery-level-checker/pkg/db"
)

type DBSrv struct {
	pb_svc_db.DBServer
	db.Database
}

func NewDBServiceServer(database *db.Database) *DBSrv {
	return &DBSrv{
		Database: *database,
	}
}
