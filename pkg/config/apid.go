package config

import "flag"

var (
	usingAuth = flag.Bool("usingAuth", true, "using firebase auth server")
	apiVersion = "v1"
	apidAddr = flag.String("apidAddr", "0.0.0.0:80", "apid address")
	usingSSL = flag.Bool("ssl", false, "using http2")
	serverCrt = flag.String("cert.crt", "/run/secrets/crt-file", "crt file location")
	serverKey = flag.String("cert.key", "/run/secrets/key-file", "ket file location")
)

type ApidConfigImpl interface {
	GetKeyPath() (string, string)
	GetApiVersion() (string)
	GetUsingSSL() (bool)
	GetUsingAuth() (bool)
	GetApidAddr() (string)
}

type ApidConfig struct {
	ApidConfigImpl
}

func (c ApidConfig) GetKeyPath() (string, string) {
	return *serverCrt, *serverKey
}

func (c ApidConfig) GetApiVersion() (string) {
	return apiVersion
}

func (c ApidConfig) GetUsingSSL() (bool) {
	return *usingSSL
}

func (c ApidConfig) GetUsingAuth() (bool) {
	return *usingAuth
}
func (c ApidConfig) GetApidAddr() (string) {
	return *apidAddr
}