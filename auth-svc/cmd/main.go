package main

import (
	"auth-svc/configs"
	"auth-svc/internal/pb"
	"auth-svc/internal/repository"
	"auth-svc/internal/service"
	"auth-svc/pkg/logging"
	"google.golang.org/grpc"
	"net"
)

func main() {
	log := logging.InitLogger()
	cfg := configs.InitConfig(log)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.WithError(err).Fatal()
	}
	//utils.FillRedis()

	//GRPC Server
	grpcServer := grpc.NewServer()

	server := &service.Server{Repo: repository.InitRepository(cfg, log)}
	pb.RegisterAuthServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.WithError(err).Fatal()
	}
}
