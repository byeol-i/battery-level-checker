package controllers

import (
	"net/http"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	firebaseSvc "github.com/byeol-i/battery-level-checker/pkg/svc/firebase"
	"go.uber.org/zap"
)

type AuthControllers struct{}

func NewAuthController() *AuthControllers {
	return &AuthControllers{}
}

func (hdl *AuthControllers) VerifyToken(next http.Handler, resp http.ResponseWriter, req *http.Request) http.Handler {
	token := req.Header.Get("Authorization")

	if len(token) < 5 {
		respondError(resp, 401, "Can't find token")
		return nil
	}

	err := firebaseSvc.CallVerifyToken(token)
	if err != nil {
		logger.Error("get some error", zap.Error(err))
		respondError(resp, 401, err.Error())
		return nil
	}

	return next
}

func (hdl *AuthControllers) ReturnServeHttp(code int, msg string) (ServeHTTP func(resp http.ResponseWriter, req *http.Request)) {
	return func(resp http.ResponseWriter, req *http.Request) {
		respondError(resp, code, msg)
	}
}

// func (hdl *AuthControllers) CreateCustom(resp http.ResponseWriter, req *http.Request) {
// 	ctx := context.Background();

// 	token, err := hdl.app.CreateCustomToken(ctx, req.Header.Get("token"))
// 	if err != nil {
// 		respondError(resp, 404, "token is not valid")
// 	}

// 	respondJSON(resp, 200, "done", token)
// }
