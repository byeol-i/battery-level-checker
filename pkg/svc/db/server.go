package dbSvc

import (
	"context"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/byeol-i/battery-level-checker/pb/unit/common"
	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
	"github.com/byeol-i/battery-level-checker/pkg/db"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/producer"
	"github.com/byeol-i/battery-level-checker/pkg/user"
)

type DBSrv struct {
	pb_svc_db.DBServer
	primaryDB db.Database
	slaveDB db.Database
	Admin *consumer.Admin
}

func NewDBServiceServer(primaryDB *db.Database, slaveDB *db.Database) *DBSrv {
	return &DBSrv{
		primaryDB: *primaryDB,
		slaveDB: *slaveDB,
		Admin: consumer.NewAdmin(),
	}
}

func (s DBSrv) AddNewUser(ctx context.Context, in *pb_svc_db.AddNewUserReq) (*pb_svc_db.AddNewUserRes, error) {
	newUser, err := user.NewUserFromProto(in.User)
	if err != nil {
		return &pb_svc_db.AddNewUserRes{Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	err = s.primaryDB.AddNewUser(newUser.UserImpl, newUser.UserCredential)
	if err != nil {
		return &pb_svc_db.AddNewUserRes{Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	admin, err := s.Admin.GetAdmin()
	if err != nil {
		logger.Error("Can't get kafka admin", zap.Error(err))
		return &pb_svc_db.AddNewUserRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	err = s.Admin.CreateTopic(admin, "battery_user__" + in.User.UserCredential.Uid)
	if err != nil {
		logger.Error("Can't create topic for device", zap.Error(err))
		return &pb_svc_db.AddNewUserRes{
			Result: &common.ReturnMsg{
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
	newDevice, err := device.NewDeviceFromProto(in.Device)
	if err != nil {
		return &pb_svc_db.AddDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	deviceId, err := s.primaryDB.AddNewDevice(*newDevice.GetDeviceSpec(), in.Uid)
	if err != nil {
		logger.Error("Can't add new device", zap.Error(err))
		return &pb_svc_db.AddDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	admin, err := s.Admin.GetAdmin()
	if err != nil {
		logger.Error("Can't get kafka admin", zap.Error(err))
		return &pb_svc_db.AddDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	err = s.Admin.CreateTopic(admin, "battery_device__" + in.Uid + "__" + deviceId)
	if err != nil {
		logger.Error("Can't create topic for device", zap.Error(err))
		return &pb_svc_db.AddDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	return &pb_svc_db.AddDeviceRes{}, nil
}

func (s DBSrv) RemoveDevice(ctx context.Context, in *pb_svc_db.RemoveDeviceReq) (*pb_svc_db.RemoveDeviceRes, error) {
	err := s.primaryDB.RemoveDevice(device.Id{
		DeviceID: in.DeviceId.Id,
	}, in.Uid.Uid)
	if err != nil {
		logger.Error("Can't remove device", zap.Error(err))
		return &pb_svc_db.RemoveDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	admin, err := s.Admin.GetAdmin()
	if err != nil {
		logger.Error("Can't get kafka admin", zap.Error(err))
		return &pb_svc_db.RemoveDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	err = s.Admin.DeleteTopic(admin, in.Uid.Uid + "__" + in.DeviceId.Id)
	if err != nil {
		logger.Error("Can't delete topic", zap.Error(err))
		return &pb_svc_db.RemoveDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}
	return &pb_svc_db.RemoveDeviceRes{}, nil
}

func (s DBSrv) GetDevices(ctx context.Context, in *pb_svc_db.GetDevicesReq) (*pb_svc_db.GetDevicesRes, error) {
	raws, err := s.slaveDB.GetDevices(in.Uid.Uid)
	if err != nil {
		return &pb_svc_db.GetDevicesRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}
	logger.Info("Get devices raws", zap.Any("raws ", raws))
	
	var devices []*pb_unit_device.Device
	for _, v := range raws {
		device := &pb_unit_device.Device{
			Spec: v.GetProtoDeviceSpec(),
		}

		devices = append(devices, device)
	}
	logger.Info("Get devices", zap.Any("devices", devices))

	return &pb_svc_db.GetDevicesRes{
		// Result: &common.ReturnMsg{
		// 	Result: "done",
		// },
		Devices: devices,
	}, nil
}

func (s DBSrv) GetBattery(ctx context.Context, in *pb_svc_db.GetBatteryReq) (*pb_svc_db.GetBatteryRes, error) {
	// raw, err := s.slaveDB.GetBattery(in.DeviceId.Id, in.Uid.Uid)
	// if err != nil {
	// 	return &pb_svc_db.GetBatteryRes{
	// 		Result: &common.ReturnMsg{
	// 			Error: err.Error(),
	// 		},
	// 	}, err
	// }
	
	// pbUnit := &pb_unit_device.BatteryLevel{
	// 	Time: timestamppb.New(*raw.Time),
	// 	BatteryLevel: int64(raw.BatteryLevel),
	// 	BatteryStatus: raw.BatteryStatus,
	// }

	// err := consumer.ConsumeLatestMessage(in.Uid.Uid+"_"+in.DeviceId.Id)
	// if err != nil {
	// 	logger.Error("Can't consume msg", zap.Error(err))
	// }
	
	return &pb_svc_db.GetBatteryRes{
		// Result: &common.ReturnMsg{
		// 	Result: pbUnit.String(),
		// },
		// BatteryLevel: pbUnit,
	}, nil
}

func (s DBSrv) GetAllBattery(ctx context.Context, in *pb_svc_db.GetAllBatteryReq) (*pb_svc_db.GetAllBatteryRes, error) {
	raws, err := s.slaveDB.GetAllBatteryLevels(in.DeviceId.Id, in.Uid.Uid)
	if err != nil {
		return &pb_svc_db.GetAllBatteryRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}
	
	var batteryLevels []*pb_unit_device.BatteryLevel
	for _, v := range raws {
		newBatteryLevel := &pb_unit_device.BatteryLevel{
			Time: timestamppb.New(*v.Time),
			BatteryLevel: int64(v.BatteryLevel),
			BatteryStatus: v.BatteryStatus,
		}

		batteryLevels = append(batteryLevels, newBatteryLevel)
	}

	return &pb_svc_db.GetAllBatteryRes{
		AllBatteryLevel: batteryLevels,
	}, nil
}

func (s DBSrv) GetUsersAllBatteryLevel(ctx context.Context, in *pb_svc_db.GetUsersAllBatteryLevelReq) (*pb_svc_db.GetUsersAllBatteryLevelRes, error) {
	raws, err := s.slaveDB.GetUsersAllBatteryLevels(in.Uid.Uid)
	if err != nil {
		return &pb_svc_db.GetUsersAllBatteryLevelRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}
	
	var batteryLevels []*pb_unit_device.BatteryLevel
	for _, v := range raws {
		newBatteryLevel := &pb_unit_device.BatteryLevel{
			Time: timestamppb.New(*v.Time),
			BatteryLevel: int64(v.BatteryLevel),
			BatteryStatus: v.BatteryStatus,
		}

		batteryLevels = append(batteryLevels, newBatteryLevel)
	}

	return &pb_svc_db.GetUsersAllBatteryLevelRes{
		Result: &common.ReturnMsg{
			Result: "Done",
		},
		AllBatteryLevel: batteryLevels,
	}, nil
}

func (s DBSrv) UpdateBatteryLevel(ctx context.Context, in *pb_svc_db.UpdateBatteryLevelReq) (*pb_svc_db.UpdateBatteryLevelRes, error) {
	batteryLevel, err := device.ProtoToBatteryLevel(in.BatteryLevel)
	if err != nil {
		return &pb_svc_db.UpdateBatteryLevelRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}
	err = producer.WriteBatteryTime(batteryLevel, in.DeviceId.Id, in.Uid.Uid)
	if err != nil {
		return &pb_svc_db.UpdateBatteryLevelRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	return &pb_svc_db.UpdateBatteryLevelRes{
		Result: &common.ReturnMsg{
			Result: "Done",
		},
	}, nil
}