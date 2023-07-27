package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	configManager := config.GetInstance()
	gRPCL, err := net.Listen("tcp", configManager.GrpcConfig.GetAuthAddr())
	if err != nil {
		return err
	}
	defer gRPCL.Close()

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	// authSrv := server.NewAuthServiceServer(firebaseApp)

	// pb_svc_firebase.RegisterFirebaseServer(grpcServer, authSrv)

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(func() error {
		logger.Info("Starting grpc server..." + configManager.GrpcConfig.GetCacheAddr())
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	return wg.Wait()
}
