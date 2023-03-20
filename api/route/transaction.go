package route

import (
	"SigmatechTechnicalTest-Golang/api/controller"
	"SigmatechTechnicalTest-Golang/api/middleware"
	"github.com/gin-gonic/gin"
)

func Transaction(r *gin.RouterGroup, c controller.Controller) {
	transactionAPI := r.Group("/transaction")
	transactionAPI.Use(middleware.VerifyToken())
	transactionAPI.POST("/buy", c.TransactionController.Buy)
	transactionAPI.GET("/list-order", c.TransactionController.ListOrder)
}
