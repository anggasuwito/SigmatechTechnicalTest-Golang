package route

import (
	"SigmatechTechnicalTest-Golang/api/controller"
	"SigmatechTechnicalTest-Golang/api/middleware"
	"github.com/gin-gonic/gin"
)

func Transaction(r *gin.RouterGroup, c controller.Controller) {
	transactionAPI := r.Group("/transaction")
	transactionAPI.Use(middleware.VerifyToken())
	transactionAPI.GET("/buy", c.TransactionController.Buy)
}
