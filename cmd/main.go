package main

import (
	"grpc-user/pkg/env"
	"grpc-user/pkg/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func init() {
	env.LoadEnvs()
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to load port :9000 %v", err)
	}
	s := user.Server{}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to load GRPC server %v", err)
	}
}
