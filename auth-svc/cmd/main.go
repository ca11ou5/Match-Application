package main

import (
	"auth-svc/configs"
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
}
