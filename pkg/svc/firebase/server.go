package server

import (
	"context"

	pb_svc_firebase "github.com/byeol-i/battery-level-checker/pb/svc/firebase"
	auth "github.com/byeol-i/battery-level-checker/pkg/authentication/firebase"
)

type AuthSrv struct {
	pb_svc_firebase.FirebaseServer
	app *auth.Firebase
}

func NewAuthServiceServer(app *auth.Firebase) *AuthSrv {
	return &AuthSrv{app: app}
}


func (s AuthSrv) AuthcationToken(ctx context.Context, in *pb_svc_firebase.AuthcationTokenReq) (*pb_svc_firebase.AuthcationTokenRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	return &pb_svc_firebase.AuthcationTokenRes{}, nil
}

func (s AuthSrv) GetUser(ctx context.Context, in *pb_svc_firebase.GetUserReq) (*pb_svc_firebase.GetUserRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	s.app.GetUser(ctx, in.Uid)

	return &pb_svc_firebase.GetUserRes{}, nil
}