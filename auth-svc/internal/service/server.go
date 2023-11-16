package service

import "auth-svc/internal/repository"

type Server struct {
	Repo *repository.Repository
}
