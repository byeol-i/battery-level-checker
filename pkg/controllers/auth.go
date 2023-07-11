package controllers

import (
	"net/http"
	"strings"

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

	const bearerPrefix = "Bearer "
	if strings.HasPrefix(token, bearerPrefix) {
		token = token[len(bearerPrefix):]

		uid, err := firebaseSvc.CallVerifyToken(token)
		if err != nil {
			logger.Error("dbSvc's error", zap.Error(err))
			respondError(resp, 401, err.Error())
			return nil
		}

		uid = strings.Replace(uid, "\"", "", -1)

		req.Header.Set("Uid", uid)
		return next
	}

	// const customTokenPrefix = "Secret "
	// if strings.HasPrefix(token, customTokenPrefix) {
	// 	token = token[len(customTokenPrefix):]

	// 	respondError(resp, 401, "Secret token not supported")
	// 	return nil
	// }

	respondError(resp, 401, "Invalid token format")
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
