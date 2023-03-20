package controller

import (
	"SigmatechTechnicalTest-Golang/api/controller/c_authentication"
	"SigmatechTechnicalTest-Golang/api/controller/c_transaction"
	"SigmatechTechnicalTest-Golang/api/usecase"
	"database/sql"
)

type Controller struct {
	AuthenticationController c_authentication.AuthenticationControllerInterface
	TransactionController    c_transaction.TransactionControllerInterface
}

func NewController(db *sql.DB) Controller {
	uc := usecase.NewUseCase(db)
	return Controller{
		AuthenticationController: c_authentication.NewAuthenticationController(uc),
		TransactionController:    c_transaction.NewTransactionController(uc),
	}
}
