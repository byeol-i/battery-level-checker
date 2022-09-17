package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	pb_svc_auth "github.com/byeol-i/battery-level-checker/pb/svc/auth"
	auth "github.com/byeol-i/battery-level-checker/pkg/auth/firebase"
	grpc_auth "github.com/byeol-i/battery-level-checker/pkg/svc/auth"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	grpcAddr = flag.String("apid grpc addr", "0.0.0.0:50010", "grpc address")
	usingTls = flag.Bool("grpc.tls", false, "using http2")
)

func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	gRPCL, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		return err
	}
	defer gRPCL.Close()

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	authSrv := grpc_auth.NewAuthServiceServer()

	pb_svc_auth.RegisterAuthServer(grpcServer, authSrv)

	firebaseApp, err := auth.InitFirebaseApp()

	auth.GetUser(context.Background(), firebaseApp)

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(func () error {
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	return wg.Wait()
}