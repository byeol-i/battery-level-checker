package dbSvc

import (
	"context"
	"flag"
	"time"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("db svc addr", "battery_db:50012", "db grpc addr")
)

func CallAddNewUser(token string) error {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_db.NewDBClient(conn)

	in := &pb_svc_db.AddNewUserReq{}

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
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
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
