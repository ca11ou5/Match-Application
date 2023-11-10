package routes

import (
	"api-gtw/internal/auth/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(ctx *gin.Context, c pb.AuthServiceClient) {
	res, err := c.SignUp(context.Background(), &pb.SignUpRequest{
		PhoneNumber: "",
		Password:    "",
		Name:        "",
		Surname:     "",
		Gender:      false,
		DateOfBirth: "",
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(int(res.Status), gin.H{"error": res.Error})
}
