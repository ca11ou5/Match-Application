package configs

import (
	"api-gtw/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port       string `env:"PORT" env-default:"8080"`
	AuthSvcUrl string `env:"AUTH_SVC_URL"`
}

func InitConfig(log *logging.Logger) *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./configs/envs/local.env", cfg)
	if err != nil {
		log.WithError(err).Fatal()
	}
	return cfg
}
