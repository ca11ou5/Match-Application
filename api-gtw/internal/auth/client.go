package auth

import (
	"api-gtw/configs"
	"api-gtw/internal/auth/pb"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	AuthClient pb.AuthServiceClient
}

func InitAuthClient(cfg *configs.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to establish connection with auth service: " + err.Error())
	}

	return pb.NewAuthServiceClient(cc)
}
