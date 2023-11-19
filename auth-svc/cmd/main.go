package main

import (
	"auth-svc/configs"
	"auth-svc/internal/pb"
	"auth-svc/internal/repository"
	"auth-svc/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := configs.InitConfig()

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal("failed to listen " + cfg.Port[1:] + ": " + err.Error())
	}
	//utils.FillRedis()

	//GRPC Server
	grpcServer := grpc.NewServer()

	server := &service.Server{Repo: repository.InitRepository(cfg)}
	pb.RegisterAuthServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve grpc server: " + err.Error())
	}
}
