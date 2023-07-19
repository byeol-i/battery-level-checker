package config

import "flag"

var (
	authAddr = flag.String("authAddr", "0.0.0.0:50010", "auth grpc address")
	usingTls = flag.Bool("grpc.tls", false, "using http2")
)

type GrpcConfig interface {
	GetAuthAddr() (string)
	GetUsingTls() (bool)
}

func GetAuthAddr() (string) {
	return *authAddr
}

func GetUsingTls() (bool) {
	return *usingTls
}