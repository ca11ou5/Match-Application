package repository

import (
	"auth-svc/configs"
	"context"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func InitPostgres(cfg *configs.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDb, cfg.PostgresSsl)
	pg, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("failed to create connection pool: " + err.Error())
	}
	err = pg.Ping(context.Background())
	if err != nil {
		fmt.Println("failed to ping: " + err.Error())
	}
	return pg
}
