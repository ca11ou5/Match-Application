package service

import (
	"auth-svc/internal/pb"
	"context"
)

func (s *Server) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return &pb.SignUpResponse{}, nil
}
