package server

import (
	pb_svc_auth "github.com/byeol-i/battery-level-checker/pb/svc/auth"
)

type AuthSrv struct {
	pb_svc_auth.AuthServer
}

func NewAuthServiceServer() *AuthSrv {
	return &AuthSrv{}
}
