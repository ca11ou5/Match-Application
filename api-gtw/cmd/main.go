package main

import (
	"api-gtw/configs"
	"api-gtw/internal/auth"
	"api-gtw/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logging.InitLogger()
	cfg := configs.InitConfig(log)

	//HTTP Server
	r := gin.Default()
	authSvc := auth.InitRoutes(r, cfg, log)
	fmt.Println(authSvc)
	err := r.Run(cfg.Port)
	if err != nil {
		log.WithError(err).Fatal()
	}
}
