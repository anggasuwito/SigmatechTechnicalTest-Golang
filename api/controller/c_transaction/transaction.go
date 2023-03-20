package c_transaction

import (
	"SigmatechTechnicalTest-Golang/api/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	UC usecase.UseCase
}

type TransactionControllerInterface interface {
	Buy(g *gin.Context)
	ListOrder(g *gin.Context)
}

func NewTransactionController(uc usecase.UseCase) TransactionControllerInterface {
	return &TransactionController{UC: uc}
}
