package main

import (
	"api-gtw/configs"
	"api-gtw/internal/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := configs.InitConfig()

	r := gin.Default()
	authSvc := auth.InitRoutes(r, cfg)
	fmt.Println(authSvc)
	err := r.Run(cfg.Port)
	if err != nil {
		log.Fatal("failed to run http server: " + err.Error())
	}

}
