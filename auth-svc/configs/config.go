package configs

import (
	"auth-svc/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port             string `env:"PORT" default-env:"50051"`
	PostgresDb       string `env:"POSTGRES_DB"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresSsl      string `env:"POSTGRES_SSL"`
}

func InitConfig(log *logging.Logger) *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./configs/envs/local.env", cfg)
	if err != nil {
		log.WithError(err).Fatal()
	}
	return cfg
}
