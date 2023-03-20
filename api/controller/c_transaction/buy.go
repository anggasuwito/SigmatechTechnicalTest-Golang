package c_transaction

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"SigmatechTechnicalTest-Golang/helper/constants"
	"SigmatechTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c TransactionController) Buy(g *gin.Context) {
	var req models.TransactionBuyRequest
	g.ShouldBindJSON(&req)
	userID, _ := g.Get(constants.UserIDKey)
	userIDString, _ := userID.(string)
	req.UserID = userIDString
	res, err := c.UC.TransactionUC.BuyUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
