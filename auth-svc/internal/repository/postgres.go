package repository

import (
	"auth-svc/configs"
	"auth-svc/pkg/logging"
	"context"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgres(cfg *configs.Config, log *logging.Logger) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDb, cfg.PostgresSsl)
	pg, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.WithError(err).Fatal()
	}
	err = pg.Ping(context.Background())
	if err != nil {
		log.WithError(err).Fatal()
	}
	return pg
}
