package c_authentication

import (
	"SigmatechTechnicalTest-Golang/api/models"
	"SigmatechTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c AuthenticationController) Login(g *gin.Context) {
	var req models.LoginRequest
	g.ShouldBind(&req)
	res, err := c.UC.AuthenticationUC.LoginUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
