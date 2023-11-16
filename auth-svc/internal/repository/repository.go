package repository

import (
	"auth-svc/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Postgres *pgxpool.Pool
}

func InitRepository(cfg *configs.Config) *Repository {
	return &Repository{Postgres: InitPostgres(cfg)}
}
