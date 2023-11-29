package auth

import (
	"api-gtw/configs"
	"api-gtw/internal/auth/routes"
	"api-gtw/pkg/logging"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, cfg *configs.Config, log *logging.Logger) *ServiceClient {
	svc := &ServiceClient{AuthClient: InitAuthClient(cfg, log)}

	auth := r.Group("auth")
	auth.POST("sign-up", svc.SignUp)
	return svc
}

func (svc *ServiceClient) SignUp(ctx *gin.Context) {
	routes.SignUp(ctx, svc.AuthClient)
}
