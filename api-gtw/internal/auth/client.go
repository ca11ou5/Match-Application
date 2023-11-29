package auth

import (
	"api-gtw/configs"
	"api-gtw/internal/auth/pb"
	"api-gtw/pkg/logging"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	AuthClient pb.AuthServiceClient
}

func InitAuthClient(cfg *configs.Config, log *logging.Logger) pb.AuthServiceClient {
	cc, err := grpc.Dial(cfg.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatal()
	}

	return pb.NewAuthServiceClient(cc)
}
