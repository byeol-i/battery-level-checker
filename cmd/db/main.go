package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb_svc_db "github.com/byeol-i/battery-level-checker/pb/svc/db"

	"github.com/byeol-i/battery-level-checker/pkg/db"
	server "github.com/byeol-i/battery-level-checker/pkg/svc/db"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	grpcAddr = flag.String("addr", "0.0.0.0:50012", "grpc address")
	usingTls = flag.Bool("tls", false, "using http2")
	test = flag.Bool("dbsvc-test", false, "testing")
)

func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	flag.Parse()

	// get env from docker, not a config pkg
	primaryDBAddr := os.Getenv("DB_PRIMARY_ADDR")
	primaryDBPort := os.Getenv("DB_PRIMARY_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")


	replicaDBAddr := os.Getenv("DB_REPLICA_ADDR")
	replicaDBPort := os.Getenv("DB_REPLICA_PORT")

	primaryDBport, err := strconv.Atoi(primaryDBPort)
	if err != nil {
		fmt.Errorf("Can't read dbPort!: %v %v", primaryDBPort, err)
		primaryDBport = 8432
	}

	replicaDBport, err := strconv.Atoi(replicaDBPort)
	if err != nil {
		fmt.Errorf("Can't read dbPort!: %v %v", replicaDBPort, err)
		primaryDBport = 8432
	}


	if *test {
		primaryDBAddr = "localhost"
		primaryDBport = 5432
		dbUser = "table_admin"
		dbPasswd = "HelloWorld"
		dbName = "battery"
	}
	

	primaryDB, err := db.ConnectDB(&db.DBConfig{
		Host:     primaryDBAddr,
		Port:     primaryDBport,
		User:     dbUser,
		Password: dbPasswd,
		DBname:   dbName,
		SSLmode:  "disable",
		// Sslmode : "verify-full",
		// Sslrootcert : "keys/ca.crt",
		// Sslkey : "keys/client.key",
		// Sslsert : "keys/client.crt",
	})
	if err != nil {
		return err
	}

	slaveDB, err := db.ConnectDB(&db.DBConfig{
		Host:     replicaDBAddr,
		Port:     replicaDBport,
		User:     dbUser,
		Password: dbPasswd,
		DBname:   dbName,
		SSLmode:  "disable",
		// Sslmode : "verify-full",
		// Sslrootcert : "keys/ca.crt",
		// Sslkey : "keys/client.key",
		// Sslsert : "keys/client.crt",
	})
	if err != nil {
		return err
	}



	gRPCL, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		return err
	}
	defer gRPCL.Close()
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	dbSrv := server.NewDBServiceServer(primaryDB, slaveDB)
	pb_svc_db.RegisterDBServer(grpcServer, dbSrv)

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(func() error {
		logger.Info("Starting grpc server..." + *grpcAddr)
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	return wg.Wait()
}
