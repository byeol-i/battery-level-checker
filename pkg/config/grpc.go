package config

import "flag"

var (
	authAddr = flag.String("authAddr", "0.0.0.0:50010", "auth grpc address")
	usingTls = flag.Bool("grpc.tls", false, "using http2")
)

type GrpcConfigImpl interface {
	GetAuthAddr() (string)
	GetUsingTls() (bool)
}
type GrpcConfig struct {
	GrpcConfigImpl
}

func (c GrpcConfig) GetAuthAddr() (string) {
	return *authAddr
}

func (c GrpcConfig) GetUsingTls() (bool) {
	return *usingTls
}