package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	AuthSvcUrl string `mapstructure:"AUTH_SVC_URL"`
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
