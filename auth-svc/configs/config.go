package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresSsl      string `mapstructure:"POSTGRES_SSL"`
}

func InitConfig() *Config {
	viper.SetConfigFile("./configs/envs/dev.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read config: " + err.Error())
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatal("failed to unmarshal config: " + err.Error())
	}

	return cfg
}
