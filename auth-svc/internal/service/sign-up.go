package service

import (
	"auth-svc/internal/pb"
	"context"
	"time"
)

func (s *Server) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	time.Sleep(1 * time.Second)
	return &pb.SignUpResponse{}, nil
}
