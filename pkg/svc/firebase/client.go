package firebaseSvc

import (
	"context"
	"flag"
	"time"

	pb_svc_firebase "github.com/byeol-i/battery-level-checker/pb/svc/firebase"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const contextTime = time.Second * 10

var (
	addr = flag.String("auth addr", "battery_auth:50010", "auth grpc addr")
)

func CallVerifyToken(token string) (string, error) {
	dialTimeout := 10 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb_svc_firebase.NewFirebaseClient(conn)

	in := &pb_svc_firebase.VerifyIdTokenReq{
		Token: token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	res, err := client.VerifyIdToken(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return "", err
	}

	return res.Result.Result, nil
}

func CallGetUser(uid string) error {
	// logger.Info("make grpc call at auth server", zap.String("uid", uid))

	dialTimeout := 3 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_firebase.NewFirebaseClient(conn)

	in := &pb_svc_firebase.GetUserReq{
		Uid: uid,
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	_, err = client.GetUser(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return err
	}

	return nil
}

func CallCreateCustomToken(token user.Token) (string, error) {
	// logger.Info("make grpc call at auth server", zap.String("uid", uid))

	dialTimeout := 3 * time.Second
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(dialTimeout))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb_svc_firebase.NewFirebaseClient(conn)

	in := &pb_svc_firebase.CreateCustomTokenReq{
		Uid: token.Uid,
		AccessToken: token.Token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()

	res, err := client.CreateCustomToken(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call")
		return "", err
	}

	return res.Token, nil
}
