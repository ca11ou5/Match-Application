package repository

import (
	"auth-svc/configs"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func InitPostgres(cfg *configs.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDb)
	pg, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("failed to create connection pool: " + err.Error())
	}

	return pg
}
