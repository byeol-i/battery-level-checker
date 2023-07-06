package db

import (
	"database/sql"
	"fmt"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Database struct {
	Conn *sql.DB
}
type DBConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	DBname      string
	SSLmode     string
	SSLrootCert string
	SSLkey      string
	SSLcert     string
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

	return &Database{Conn: db, }, nil
}

func ErrorHandlingMsg(err error) (error) {
	if err, ok := err.(*pq.Error); ok {
		logger.Error("pq Error!", zap.Error(err), zap.Any("pq code",err.Code))
		switch {
		// case errors.As(err, *pq.Error.)
		default:
			return err
		}
	}

	return err
}