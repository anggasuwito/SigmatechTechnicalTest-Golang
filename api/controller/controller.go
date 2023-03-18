package controller

import (
	"SigmatechTechnicalTest-Golang/api/controller/c_authentication"
	"SigmatechTechnicalTest-Golang/api/usecase"
	"database/sql"
)

type Controller struct {
	AuthenticationController c_authentication.AuthenticationControllerInterface
}

func NewController(db *sql.DB) Controller {
	uc := usecase.NewUseCase(db)
	return Controller{
		AuthenticationController: c_authentication.NewAuthenticationController(uc),
	}
}
