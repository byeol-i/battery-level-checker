package firebaseSvc

import (
	"context"
	"flag"
	"time"

	pb_svc_firebase "github.com/byeol-i/battery-level-checker/pb/svc/firebase"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("auth addr", "app_auth:50010", "auth grpc addr")
)

func CallVerifyToken(token string) error {
	// logger.Info("make grpc call at auth server", zap.String("token", token))

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_firebase.NewFirebaseClient(conn)

	in := &pb_svc_firebase.VerifyTokenReq{
		Token: token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_, err = client.VerifyToken(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}

func CallGetUser(uid string) error {
	// logger.Info("make grpc call at auth server", zap.String("uid", uid))

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_firebase.NewFirebaseClient(conn)

	in := &pb_svc_firebase.GetUserReq{
		Uid: uid,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_, err = client.GetUser(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}	