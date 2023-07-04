package firebaseSvc

import (
	"context"
	"encoding/json"

	pb_svc_firebase "github.com/byeol-i/battery-level-checker/pb/svc/firebase"
	pb_unit_common "github.com/byeol-i/battery-level-checker/pb/unit/common"
	auth "github.com/byeol-i/battery-level-checker/pkg/authentication/firebase"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
)

type AuthSrv struct {
	pb_svc_firebase.FirebaseServer
	app *auth.FirebaseApp
}

func NewAuthServiceServer(app *auth.FirebaseApp) *AuthSrv {
	return &AuthSrv{app: app}
}

func (s AuthSrv) GetUser(ctx context.Context, in *pb_svc_firebase.GetUserReq) (*pb_svc_firebase.GetUserRes, error) {
	result, err := s.app.GetUser(ctx, in.Uid)
	if err != nil {
		return &pb_svc_firebase.GetUserRes{
			// Result: "",
			Result: &pb_unit_common.ReturnMsg{
				Error: err.Error(),
			},	
		}, err
	}

	jsonRes, err := json.Marshal(result)
	if err != nil {
		return &pb_svc_firebase.GetUserRes{
			// Result: "",
			Result: &pb_unit_common.ReturnMsg{
				Error: err.Error(),
			},	
		}, err
	}

	return &pb_svc_firebase.GetUserRes{
		Result: &pb_unit_common.ReturnMsg{
			Result: string(jsonRes),
		},	
	}, nil
}

func (s AuthSrv) CreateCustomToken(ctx context.Context, in *pb_svc_firebase.CreateCustomTokenReq) (*pb_svc_firebase.CreateCustomTokenRes, error) {
	token, err := s.app.CreateCustomToken(ctx, in.Uid)
	if err != nil {
		return &pb_svc_firebase.CreateCustomTokenRes{}, err
	}

	return &pb_svc_firebase.CreateCustomTokenRes{
		Token: token,
	}, nil
}

type Res struct {
	Uid string `json:"uid"`
}

func (s AuthSrv) VerifyIdToken(ctx context.Context, in *pb_svc_firebase.VerifyIdTokenReq) (*pb_svc_firebase.VerifyIdTokenRes, error) {
	result, err := s.app.VerifyIDToken(ctx, in.Token)
	if err != nil {
		logger.Error("Can't verify token", zap.Error(err))
		return &pb_svc_firebase.VerifyIdTokenRes{
			Result: &pb_unit_common.ReturnMsg{
				Error: err.Error(),
			},	
		}, err
	}
	
	jsonRes, err := json.Marshal(result)
	if err != nil {
		return &pb_svc_firebase.VerifyIdTokenRes{
			// Result: "",
			Result: &pb_unit_common.ReturnMsg{
				Error: err.Error(),
			},	
		}, err
	}

	return &pb_svc_firebase.VerifyIdTokenRes{
		Result: &pb_unit_common.ReturnMsg{
			Result: string(jsonRes),
		},	
		// Error:  "",
	}, nil
}

func (s AuthSrv) GetUserIdByIdToken(ctx context.Context, in *pb_svc_firebase.GetUserIdByIdTokenReq) (*pb_svc_firebase.GetUserIdByIdTokenRes, error) {
	result, err := s.app.VerifyIDToken(ctx, in.Token)
	if err != nil {
		logger.Error("Can't verify token", zap.Error(err))
		return &pb_svc_firebase.GetUserIdByIdTokenRes{
			Result: &pb_unit_common.ReturnMsg{
				Error: err.Error(),
			},	
		}, err
	}

	jsonRes, err := json.Marshal(result)
	if err != nil {
		return &pb_svc_firebase.GetUserIdByIdTokenRes{
			// Result: "",
			Result: &pb_unit_common.ReturnMsg{
				Error: err.Error(),
			},	
		}, err
	}

	res := Res{}
	err = json.Unmarshal([]byte(jsonRes), &res)
	if err != nil {
		return &pb_svc_firebase.GetUserIdByIdTokenRes{
			// Result: "",
			Result: &pb_unit_common.ReturnMsg{
				Error: err.Error(),
			},	
		}, err
	}

	return &pb_svc_firebase.GetUserIdByIdTokenRes{
		Result: &pb_unit_common.ReturnMsg{
			Result: res.Uid,
		},	
		// Error:  "",
	}, nil
}