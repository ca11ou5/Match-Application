package routes

import (
	"api-gtw/internal/auth/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpRequestBody struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname"`
	Gender      bool   `json:"gender"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	City        string `json:"city" binding:"required"`
}

func SignUp(ctx *gin.Context, c pb.AuthServiceClient) {
	b := SignUpRequestBody{}
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.SignUp(context.Background(), &pb.SignUpRequest{
		PhoneNumber: "",
		Password:    "",
		Name:        "",
		Surname:     "",
		Gender:      false,
		DateOfBirth: "",
		City:        "",
	})
	if err != nil {
		//ctx.AbortWithError(http.StatusBadGateway, err)
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(int(res.Status), gin.H{"error": res.Error})
}
