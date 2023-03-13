package dbSvc

import (
	"context"
	"flag"
	"time"

	pb_unit_device "github.com/byeol-i/battery-level-checker/pb/unit/device"
	pb_unit_user "github.com/byeol-i/battery-level-checker/pb/unit/user"
	"go.uber.org/zap"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("dbsvc-addr", "battery_db:50012", "db grpc addr")
)

func CallAddNewUser(userSpec *user.UserImpl) error {
	dialTimeout := 5 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(),  grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.AddNewUserReq{
		User: &pb_unit_user.User{
			Id: &pb_unit_user.ID{
				Id: userSpec.Id,
			},
			Name: userSpec.Name,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_, err = client.AddNewUser(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}

func CallAddNewDevice(newDevice *device.Device) error {
	dialTimeout := 5 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	pbUnit := newDevice.ToProtoDevice()

	in := &pb_svc_db.AddDeviceReq{
		Device: pbUnit,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_, err = client.AddDevice(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}

func CallRemoveDevice(deviceID string, uid string) error {
	dialTimeout := 5 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.RemoveDeviceReq{
		
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_, err = client.RemoveDevice(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}

func CallGetAllBattery(deviceID string, uid string) ([]*device.BatteryLevel, error) {
	dialTimeout := 5 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.GetAllBatteryReq{
		DeviceId: &pb_unit_device.ID{
			Id: deviceID,
		},
		Uid: &pb_unit_user.ID{
			Id: uid,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

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
	dialTimeout := 5 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.GetBatteryReq{
		DeviceId: &pb_unit_device.ID{
			Id: deviceID,
		},
		Uid: &pb_unit_user.ID{
			Id: uid,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	res, err := client.GetBattery(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return nil, err
	}

	newBatteryLevel, err := device.ProtoToBatteryLevel(res.BatteryLevel)
	if err != nil {
		logger.Error("Can't make pb to batteryLevel struct", zap.Error(err))
		return nil, err
	}

	return newBatteryLevel, nil
}

func CallUpdateBatteryLevel(deviceID string, uid string, batteryLevel *device.BatteryLevel) error {
	dialTimeout := 5 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	// device, err := device

	in := &pb_svc_db.UpdateBatteryLevelReq{
		BatteryLevel: &pb_unit_device.BatteryLevel{
			Time: batteryLevel.Time.String(),
			BatteryLevel: int64(batteryLevel.BatteryLevel),
			BatteryStatus: batteryLevel.BatteryStatus,
		},
		DeviceId: &pb_unit_device.ID{
			Id: deviceID,
		},
		Uid: &pb_unit_user.ID{
			Id: uid,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_, err = client.UpdateBatteryLevel(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}