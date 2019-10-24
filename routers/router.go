package routers

import (
	"github.com/jinweida/fish/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/assets", "./assets")
	// reg router
	api := r.Group("/api/")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	wxpay := &controller.WXPayController{}
	api.GET("/wxpay/test", wxpay.Test)
	return r
}