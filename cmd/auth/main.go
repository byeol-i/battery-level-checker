package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	pb_svc_firebase "github.com/byeol-i/battery-level-checker/pb/svc/firebase"
	auth "github.com/byeol-i/battery-level-checker/pkg/authentication/firebase"
	server "github.com/byeol-i/battery-level-checker/pkg/svc/firebase"

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

	firebaseApp, err := auth.NewFirebaseApp()
	
	authSrv := server.NewAuthServiceServer(firebaseApp)

	pb_svc_firebase.RegisterFirebaseServer(grpcServer, authSrv)

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