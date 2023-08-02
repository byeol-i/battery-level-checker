package config

import "flag"

var (
	authAddr = flag.String("authAddr", "0.0.0.0:50010", "auth grpc address")
	cacheAddr = flag.String("cacheAddr", "0.0.0.0:50015", "cache grpc address")
	dbAddr = flag.String("dbAddr", "0.0.0.0:50012", "db grpc address")
	dbSvcAddr = flag.String("dbSvcAddr", "battery_db:50012", "db svc grpc addr")
	cacheSvcAddr = flag.String("cacheSvcAddr", "battery_cache:50015", "cache svc grpc addr")

	usingTls = flag.Bool("grpc.tls", false, "using http2")
)

type GrpcConfigImpl interface {
	GetAuthAddr() (string)
	GetCacheAddr() (string)
	GetCacheSvcAddr() (string)
	GetDBAddr() (string)
	GetDBSvcAddr() (string)
	GetUsingTls() (bool)
}
type GrpcConfig struct {
	GrpcConfigImpl
}

func (c GrpcConfig) GetDBSvcAddr() (string) {
	return *dbSvcAddr
}

func (c GrpcConfig) GetCacheSvcAddr() (string) {
	return *cacheSvcAddr
}

func (c GrpcConfig) GetDBAddr() (string) {
	return *dbAddr
}

func (c GrpcConfig) GetAuthAddr() (string) {
	return *authAddr
}

func (c GrpcConfig) GetCacheAddr() (string) {
	return *cacheAddr
}

func (c GrpcConfig) GetUsingTls() (bool) {
	return *usingTls
}