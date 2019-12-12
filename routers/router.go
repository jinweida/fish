package routers

import (
	"fish.com/fish-front-api/controller"
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

	login:=&controller.LoginController{}
	api.GET("/login",login.Login)

	wxpay := &controller.WXPayController{}
	api.GET("/wxpay/test", wxpay.Test)
	//api.Use(middleware.JWTAuth())

	sysConfig:=&controller.SysConfigController{}
	api.GET("/sysconfig",sysConfig.List)

	sysAdvert:=&controller.SysAdvertController{}
	api.GET("/sysadvert",sysAdvert.List)

	merchantClassify:=&controller.MerchantClassifyController{}
	api.GET("/merchant/classify",merchantClassify.List)


	merchantStore:=&controller.MerchantStoreController{}
	api.GET("/merchant/store",merchantStore.List)
	return r
}