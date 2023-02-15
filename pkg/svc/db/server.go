package db

import (
	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
)

type DBSrv struct {
	pb_svc_db.DBServer
}

func NewDBServiceServer() *DBSrv {
	return &DBSrv{}
}
