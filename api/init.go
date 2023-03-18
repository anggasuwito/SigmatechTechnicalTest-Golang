package api

import (
	"SigmatechTechnicalTest-Golang/api/controller"
	"SigmatechTechnicalTest-Golang/api/route"
	"SigmatechTechnicalTest-Golang/config"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()
	r.GET("", func(context *gin.Context) { return })
	v1Api := r.Group("/v1")
	c := controller.NewController(config.ConnDB())

	route.Authentication(v1Api, c)

	r.Run(os.Getenv("ADDRESS"))
}
