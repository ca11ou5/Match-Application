package repository

import (
	"auth-svc/configs"
	"auth-svc/pkg/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Postgres *pgxpool.Pool
}

func InitRepository(cfg *configs.Config, log *logging.Logger) *Repository {
	return &Repository{Postgres: InitPostgres(cfg, log)}
}
