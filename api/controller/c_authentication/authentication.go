package c_authentication

import (
	"SigmatechTechnicalTest-Golang/api/usecase"
	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	UC usecase.UseCase
}

type AuthenticationControllerInterface interface {
	Login(g *gin.Context)
}

func NewAuthenticationController(uc usecase.UseCase) AuthenticationControllerInterface {
	return &AuthenticationController{UC: uc}
}
