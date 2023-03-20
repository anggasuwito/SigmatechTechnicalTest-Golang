package c_transaction

import (
	"SigmatechTechnicalTest-Golang/helper/constants"
	"SigmatechTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c TransactionController) ListOrder(g *gin.Context) {
	userID, _ := g.Get(constants.UserIDKey)
	userIDString, _ := userID.(string)
	res, err := c.UC.TransactionUC.ListOrderUC(userIDString)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
