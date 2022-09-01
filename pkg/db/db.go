package db

import (
	"database/sql"
	"fmt"
)

type Database struct {
	Conn *sql.DB
}
type DBConfig struct {
	Host string 
	Port int 
	User string 
	Password string 
	DBname string 
	SSLmode string 
	SSLrootCert string 
	SSLkey string 
	SSLcert string
}

func ConnectDB(config *DBConfig) (*Database, error) {
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	config.Host, config.Port, config.User, config.Password, config.DBname, config.SSLmode)
	
	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	return &Database{Conn: db}, nil
}

