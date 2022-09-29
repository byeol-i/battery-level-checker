package firebaseSvc

import (
	"context"

	pb_svc_firebase "github.com/byeol-i/battery-level-checker/pb/svc/firebase"
	auth "github.com/byeol-i/battery-level-checker/pkg/authentication/firebase"
)

type AuthSrv struct {
	pb_svc_firebase.FirebaseServer
	app *auth.FirebaseApp
}

func NewAuthServiceServer(app *auth.FirebaseApp) *AuthSrv {
	return &AuthSrv{app: app}
}

func (s AuthSrv) GetUser(ctx context.Context, in *pb_svc_firebase.GetUserReq) (*pb_svc_firebase.GetUserRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	_, err := s.app.GetUser(ctx, in.Uid)
	if err != nil {
		return &pb_svc_firebase.GetUserRes{}, err
	}

	return &pb_svc_firebase.GetUserRes{}, nil
}

func (s AuthSrv) CreateCustomToken(ctx context.Context, in *pb_svc_firebase.CreateCustomTokenReq) (*pb_svc_firebase.CreateCustomTokenRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }

	token, err := s.app.CreateCustomToken(ctx, in.Uid)
	if err != nil {
		return &pb_svc_firebase.CreateCustomTokenRes{}, err
	}

	return &pb_svc_firebase.CreateCustomTokenRes{
		Token: token,
	}, nil
}

func (s AuthSrv) VerifyToken(ctx context.Context, in *pb_svc_firebase.VerifyTokenReq) (*pb_svc_firebase.VerifyTokenRes, error) {
	// if in != nil {
	// 	logger.Error("in is not nil")
	// }
	result, err := s.app.VerifyIDToken(ctx, in.Token)
	if err != nil {
		return &pb_svc_firebase.VerifyTokenRes{}, err
	}
	

	return &pb_svc_firebase.VerifyTokenRes{
		Msg: result,
	}, nil
}