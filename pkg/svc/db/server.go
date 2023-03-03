package dbSvc

import (
	"context"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"

	"github.com/byeol-i/battery-level-checker/pkg/db"
)

type DBSrv struct {
	pb_svc_db.DBServer
	db db.Database
}

func NewDBServiceServer(database *db.Database) *DBSrv {
	return &DBSrv{
		db: *database,
	}
}

func (s DBSrv) SignUp(ctx context.Context, in *pb_svc_db.SignUpReq) (*pb_svc_db.SignUpRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	return &pb_svc_db.SignUpRes{}, nil
}

func (s DBSrv) AddDevice(ctx context.Context, in *pb_svc_db.AddDeviceReq) (*pb_svc_db.AddDeviceRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	return &pb_svc_db.AddDeviceRes{}, nil
}

func (s DBSrv) RemoveDevice(ctx context.Context, in *pb_svc_db.RemoveDeviceReq) (*pb_svc_db.RemoveDeviceRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	return &pb_svc_db.RemoveDeviceRes{}, nil
}

func (s DBSrv) GetDevices(ctx context.Context, in *pb_svc_db.GetDevicesReq) (*pb_svc_db.GetDevicesRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	// raws, err := s.db.GetDevices(in.Id)
	// if err != nil {
	// 	return &pb_svc_db.GetDevicesRes{
	// 		Error: err.Error(),
	// 	}, err
	// }

	// var devices []*pb_unit_device.Device
	// for _, v := range *raws {
	// 	device := &pb_unit_device.Device{
	// 		Spec: v.GetDeviceSpec().ToProtoModel(),
	// 	}

	// }


	return &pb_svc_db.GetDevicesRes{
		// Devices: devices,
	}, nil
}
