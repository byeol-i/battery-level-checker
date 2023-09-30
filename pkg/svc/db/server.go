package dbSvc

import (
	"context"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/byeol-i/battery-level-checker/pb/unit/common"
	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
	"github.com/byeol-i/battery-level-checker/pkg/db"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/topic"
	"github.com/byeol-i/battery-level-checker/pkg/user"
)

type DBSrv struct {
	pb_svc_db.DBServer
	primaryDB db.Database
	slaveDB db.Database
	TopicManager *topic.TopicManager
}

func NewDBServiceServer(primaryDB *db.Database, slaveDB *db.Database) *DBSrv {
	return &DBSrv{
		primaryDB: *primaryDB,
		slaveDB: *slaveDB,
		TopicManager: topic.NewTopicManager(),
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

	// admin, err := s.TopicManager.GetAdmin()
	// if err != nil {
	// 	logger.Error("Can't get kafka admin", zap.Error(err))
	// 	return &pb_svc_db.AddNewUserRes{
	// 		Result: &common.ReturnMsg{
	// 			Error: err.Error(),
	// 		},
	// 	}, err
	// }

	// err = s.TopicManager.CreateTopic(admin, "battery_user_" + in.User.UserCredential.Uid)
	// if err != nil {
	// 	logger.Error("Can't create topic for device", zap.Error(err))
	// 	return &pb_svc_db.AddNewUserRes{
	// 		Result: &common.ReturnMsg{
	// 			Error: err.Error(),
	// 		},
	// 	}, err
	// }

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

	_, err = s.primaryDB.AddNewDevice(*newDevice.GetDeviceSpec(), in.Uid)
	if err != nil {
		logger.Error("Can't add new device", zap.Error(err))
		return &pb_svc_db.AddDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	// admin, err := s.TopicManager.GetAdmin()
	// if err != nil {
	// 	logger.Error("Can't get kafka admin", zap.Error(err))
	// 	return &pb_svc_db.AddDeviceRes{
	// 		Result: &common.ReturnMsg{
	// 			Error: err.Error(),
	// 		},
	// 	}, err
	// }

	// err = s.TopicManager.CreateTopic(admin, "battery_device_" + in.Uid + "_" + deviceId)
	// if err != nil {
	// 	logger.Error("Can't create topic for device", zap.Error(err))
	// 	return &pb_svc_db.AddDeviceRes{
	// 		Result: &common.ReturnMsg{
	// 			Error: err.Error(),
	// 		},
	// 	}, err
	// }

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

	admin, err := s.TopicManager.GetAdmin()
	if err != nil {
		logger.Error("Can't get kafka admin", zap.Error(err))
		return &pb_svc_db.RemoveDeviceRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}

	err = s.TopicManager.DeleteTopic(admin, in.Uid.Uid + "_" + in.DeviceId.Id)
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
	
	var devices []*pb_unit_device.Device
	for _, v := range raws {
		device := &pb_unit_device.Device{
			Spec: v.GetProtoDeviceSpec(),
			Id: &pb_unit_device.ID{
				Id: v.GetDeviceId(),
			},
		}

		devices = append(devices, device)
	}

	return &pb_svc_db.GetDevicesRes{
		Devices: devices,
	}, nil
}

func (s DBSrv) GetBattery(ctx context.Context, in *pb_svc_db.GetBatteryReq) (*pb_svc_db.GetBatteryRes, error) {
	raw, err := s.slaveDB.GetBattery(in.DeviceId.Id, in.Uid.Uid)
	if err != nil {
		return &pb_svc_db.GetBatteryRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err
	}
	
	pbUnit := &pb_unit_device.BatteryLevel{
		Time: timestamppb.New(raw.Time.Time),
		BatteryLevel: int64(raw.BatteryLevel),
		BatteryStatus: raw.BatteryStatus,
	}
	
	return &pb_svc_db.GetBatteryRes{
		Result: &common.ReturnMsg{
			Result: pbUnit.String(),
		},
		BatteryLevel: pbUnit,
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
			Time: timestamppb.New(v.Time.Time),
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
			Time: timestamppb.New(v.Time.Time),
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
	
	err = s.primaryDB.UpdateBattery(in.DeviceId.Id, in.Uid.Uid, batteryLevel)
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