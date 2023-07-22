package dbSvc

import (
	"context"
	"flag"
	"time"

	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
	pb_unit_user "github.com/byeol-i/battery-level-checker/pb/unit/user"
	"go.uber.org/zap"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/producer"
	"github.com/byeol-i/battery-level-checker/pkg/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const contextTime = time.Second * 5

var (
	addr = flag.String("dbsvc-addr", "battery_db:50012", "db grpc addr")
)

func CallAddNewUser(userSpec *user.UserImpl, userCredential *user.UserCredential) error {	
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()
	
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(),  grpc.WithTimeout(dialTimeout))
	
	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.AddNewUserReq{
		User: &pb_unit_user.User{
			UserCredential: &pb_unit_user.UserCredential{
				Uid: userCredential.Uid,
			},
			Name: userSpec.Name,
			Email: userSpec.Email,
		},
	}

	_, err = client.AddNewUser(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call", zap.Error(err) )
		return err
	}

	return nil
}

func CallAddNewDevice(newDevice *device.Device, uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()
	
	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	pbUnit := newDevice.ToProtoDevice()

	in := &pb_svc_db.AddDeviceReq{
		Device: pbUnit,
		Uid: uid,
	}

	_, err = client.AddDevice(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}

func CallRemoveDevice(deviceID string, uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.RemoveDeviceReq{
		DeviceId: &pb_unit_device.ID{
			Id: uid,
		},
	}

	_, err = client.RemoveDevice(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}

func CallGetUsersAllBattery(uid string) ([]*device.BatteryLevel, error) {	
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.GetUsersAllBatteryLevelReq{
		Uid: &pb_unit_user.UserCredential{
			Uid: uid,
		},
	}

	res, err := client.GetUsersAllBatteryLevel(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return nil, err
	}
	
	allBatteryLevel := []*device.BatteryLevel{}
	for _, v := range res.AllBatteryLevel {
		newBatteryLevel, err := device.ProtoToBatteryLevel(v)
		if err != nil {
			logger.Error("Can't make pb to batteryLevel struct", zap.Error(err))
		}

		allBatteryLevel = append(allBatteryLevel, newBatteryLevel)
	}

	return allBatteryLevel, nil
}

func CallGetAllDevices(uid string) ([]*device.Device, error) {	
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()
	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)
	in := &pb_svc_db.GetDevicesReq{
		Uid: &pb_unit_user.UserCredential{
			Uid: uid,
		},
	}

	res, err := client.GetDevices(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return nil, err
	}
	allDevices := []*device.Device{}

	for _, v := range res.Devices {
		newDevice, err := device.NewDeviceFromProto(v)
		if err != nil {
			logger.Error("Can't make pb to device struct", zap.Error(err))
		}

		allDevices = append(allDevices, newDevice)
	}

	return allDevices, nil
}

func CallGetAllBattery(deviceID string, uid string) ([]*device.BatteryLevel, error) {	
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.GetAllBatteryReq{
		DeviceId: &pb_unit_device.ID{
			Id: deviceID,
		},
		Uid: &pb_unit_user.UserCredential{
			Uid: uid,
		},
	}

	res, err := client.GetAllBattery(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return nil, err
	}
	
	allBatteryLevel := []*device.BatteryLevel{}
	for _, v := range res.AllBatteryLevel {
		newBatteryLevel, err := device.ProtoToBatteryLevel(v)
		if err != nil {
			logger.Error("Can't make pb to batteryLevel struct", zap.Error(err))
		}

		allBatteryLevel = append(allBatteryLevel, newBatteryLevel)
	}

	return allBatteryLevel, nil
}

func CallGetBattery(deviceID string, uid string) (*device.BatteryLevel, error) {	
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()
	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// client := pb_svc_db.NewDBClient(conn)

	// in := &pb_svc_db.GetBatteryReq{
	// 	DeviceId: &pb_unit_device.ID{
	// 		Id: deviceID,
	// 	},
	// 	Uid: &pb_unit_user.UserCredential{
	// 		Uid: uid,
	// 	},
	// }


	err = consumer.ConsumeLatestMessage("battery_user__" + uid + "_" + deviceID)
	if err != nil {
		return nil, err
	}

	return nil, nil

	// res, err := client.GetBattery(ctx, in)
	// if err != nil {
	// 	logger.Error("Can't call grpc call")
	// 	return nil, err
	// }

	// newBatteryLevel, err := device.ProtoToBatteryLevel(res.BatteryLevel)
	// if err != nil {
	// 	logger.Error("Can't make pb to batteryLevel struct", zap.Error(err))
	// 	return nil, err
	// }

	// return newBatteryLevel, nil
}

func CallUpdateBatteryLevel(deviceID string, uid string, batteryLevel *device.BatteryLevel) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.UpdateBatteryLevelReq{
		BatteryLevel: &pb_unit_device.BatteryLevel{
			Time: timestamppb.New(*batteryLevel.Time),
			BatteryLevel: int64(batteryLevel.BatteryLevel),
			BatteryStatus: batteryLevel.BatteryStatus,
		},
		DeviceId: &pb_unit_device.ID{
			Id: deviceID,
		},
		Uid: &pb_unit_user.UserCredential{
			Uid: uid,
		},
	}

	_, err = client.UpdateBatteryLevel(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call", zap.Error(err))
		return err
	}

	err = producer.WriteBatteryTime(batteryLevel, deviceID, uid)
	if err != nil {
		logger.Error("Can't producer msg", zap.Error(err))
		return err
	}

	return nil
}