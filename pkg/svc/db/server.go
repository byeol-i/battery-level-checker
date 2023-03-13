package dbSvc

import (
	"context"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
	"go.uber.org/zap"

	"github.com/byeol-i/battery-level-checker/pb/unit/common"
	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
	"github.com/byeol-i/battery-level-checker/pkg/db"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/user"
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

func (s DBSrv) AddNewUser(ctx context.Context, in *pb_svc_db.AddNewUserReq) (*pb_svc_db.AddNewUserRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	newUser, err := user.NewUserFromProto(in.User)
	if err != nil {
		return &pb_svc_db.AddNewUserRes{Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	err = s.db.AddNewUser(newUser.UserImpl)
	if err != nil {
		return &pb_svc_db.AddNewUserRes{Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	return &pb_svc_db.AddNewUserRes{
		Result: &common.ReturnMsg{
			Result: "success",
		},
	}, nil
}

func (s DBSrv) AddDevice(ctx context.Context, in *pb_svc_db.AddDeviceReq) (*pb_svc_db.AddDeviceRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	newDevice, err := device.NewDeviceFromProto(in.Device)
	if err != nil {
		return &pb_svc_db.AddDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	err = s.db.AddNewDevice(*newDevice.GetDeviceSpec(), in.Uid)
	if err != nil {
		logger.Error("Can't add new device", zap.Error(err))
		return &pb_svc_db.AddDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	return &pb_svc_db.AddDeviceRes{}, nil
}

func (s DBSrv) RemoveDevice(ctx context.Context, in *pb_svc_db.RemoveDeviceReq) (*pb_svc_db.RemoveDeviceRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	err := s.db.RemoveDevice(device.Id{
		DeviceID: in.Uid.Id,
	})
	if err != nil {
		logger.Error("Can't remove device", zap.Error(err))
		return &pb_svc_db.RemoveDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	return &pb_svc_db.RemoveDeviceRes{}, nil
}

func (s DBSrv) GetDevices(ctx context.Context, in *pb_svc_db.GetDevicesReq) (*pb_svc_db.GetDevicesRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	raws, err := s.db.GetDevices(in.Uid.Id)
	if err != nil {
		return &pb_svc_db.GetDevicesRes{
			Error: err.Error(),
		}, err
	}
	
	var devices []*pb_unit_device.Device
	for _, v := range raws {
		device := &pb_unit_device.Device{
			Spec: v.GetProtoDeviceSpec(),
		}

		devices = append(devices, device)
	}

	return &pb_svc_db.GetDevicesRes{
		Devices: devices,
	}, nil
}


func (s DBSrv) GetBattery(ctx context.Context, in *pb_svc_db.GetBatteryReq) (*pb_svc_db.GetBatteryRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	// raws, err := s.db.GetDevices(in.Uid)
	// if err != nil {
	// 	return &pb_svc_db.GetBatteryRes{
	// 		Error: err.Error(),
	// 	}, err
	// }
	
	// var devices []*pb_unit_device.Device
	// for _, v := range raws {
	// 	device := &pb_unit_device.Device{
	// 		Spec: v.GetProtoDeviceSpec(),
	// 	}

	// 	devices = append(devices, device)
	// }

	return &pb_svc_db.GetBatteryRes{
	}, nil
}
