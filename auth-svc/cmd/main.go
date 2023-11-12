package main

import (
	"auth-svc/configs"
	"auth-svc/internal/pb"
	"auth-svc/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := configs.InitConfig()

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal("failed to listen ")
	}

	//GRPC Server
	grpcServer := grpc.NewServer()
	server := service.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve grpc server: " + err.Error())
	}
}
