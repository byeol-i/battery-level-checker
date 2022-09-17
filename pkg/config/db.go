package config

import (
	"github.com/byeol-i/battery-level-checker/pkg/db"
	"gopkg.in/alecthomas/kingpin.v2"
)


var (
	dbHost			= kingpin.Flag("dbHost", "db's addr").Default("localhost").String()
	dbPort			= kingpin.Flag("dbPort", "db's port").Default("5432").Int()
	dbUser			= kingpin.Flag("dbUser", "db's user id").Default("table_admin").String()
	dbPassword		= kingpin.Flag("dbPassword", "db's user password").Default("HelloWorld").String()
	dbName			= kingpin.Flag("dbName", "db's name").Default("kafka").String()
	dbSSLmode		= kingpin.Flag("dbSSLMode", "Retry limit").Default("disable").String()
	dbSSLrootCert	= kingpin.Flag("dbSSLrootCert", "db's rootCert").Default("").String()
	dbSSLkey		= kingpin.Flag("dbSSLkey", "db's SSLkey").Default("").String()
	dbSSLcert		= kingpin.Flag("dbSSLcert", "db's SSLcert").Default("").String()
)

func GetDBConfig() *db.DBConfig {
	kingpin.Parse()
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