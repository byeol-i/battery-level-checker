package client

import (
	"context"
	"flag"
	"time"

	pb_svc_firebase "github.com/byeol-i/battery-level-checker/pb/svc/firebase"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("auth addr", "app_auth:50010", "auth grpc addr")
)

func CallAuth(token string) error {
	logger.Info("make grpc call at auth server", zap.String("token", token))

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_firebase.NewFirebaseClient(conn)

	in := &pb_svc_firebase.AuthcationTokenReq{
		Token: token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = client.AuthcationToken(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}	