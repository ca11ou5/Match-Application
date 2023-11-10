package auth

import (
	"api-gtw/configs"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, cfg *configs.Config) *ServiceClient {
	svc := &ServiceClient{AuthClient: InitAuthClient(cfg)}

	auth := r.Group("auth")
	auth.POST("sign-up")
	return svc
}
