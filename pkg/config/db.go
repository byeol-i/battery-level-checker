package config

import (
	"flag"

	"github.com/byeol-i/battery-level-checker/pkg/db"
)

var (
	dbHost			= flag.String("dbHost", "localhost", "db's addr")
	dbPort			= flag.Int("dbPort", 5432, "db's port")
	dbUser			= flag.String("dbUser", "table_admin", "db's user id")
	dbPassword		= flag.String("dbPassword", "HelloWorld", "db's user password")
	dbName			= flag.String("dbName", "battery", "db's name")
	dbSSLmode		= flag.String("dbSSLMode", "disable", "Retry limit")
	dbSSLrootCert	= flag.String("dbSSLrootCert", "", "db's rootCert")
	dbSSLkey		= flag.String("dbSSLkey", "", "db's SSLkey")
	dbSSLcert		= flag.String("dbSSLcert", "", "db's SSLcert")
)

type DBConfigImpl interface {
	GetDBConfig() *db.DBConfig
}

type DBConfig struct {
	DBConfigImpl
}

func GetDBConfig() *db.DBConfig {
	return &db.DBConfig{
		Host: *dbHost,
		Port: *dbPort,
		User: *dbUser,
		Password: *dbPassword,
		DBname: *dbName,
		SSLmode: *dbSSLmode,
		SSLrootCert: *dbSSLrootCert,
		SSLkey: *dbSSLkey,
		SSLcert: *dbSSLcert,
	}
}