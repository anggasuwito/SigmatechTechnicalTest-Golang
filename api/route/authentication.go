package route

import (
	"SigmatechTechnicalTest-Golang/api/controller"
	"github.com/gin-gonic/gin"
)

func Authentication(r *gin.RouterGroup, c controller.Controller) {
	orderAPI := r.Group("/auth")
	orderAPI.GET("/login", c.AuthenticationController.Login)
}
