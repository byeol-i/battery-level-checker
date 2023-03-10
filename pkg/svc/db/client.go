package dbSvc

import (
	"context"
	"flag"
	"time"

	pb_unit_user "github.com/byeol-i/battery-level-checker/pb/unit/user"

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
	dialTimeout := 3 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(),  grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.AddNewUserReq{
		User: &pb_unit_user.User{
			Id: &pb_unit_user.ID{
				Uuid: userSpec.Id,
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
	dialTimeout := 3 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	pbUnit, err := newDevice.ToProto()
	if err != nil {
		logger.Error("Can't translate to pb unit")
		return err
	}

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

func CallRemoveDevice(deviceID string) error {
	dialTimeout := 3 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.RemoveDeviceReq{
		Id: deviceID,
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