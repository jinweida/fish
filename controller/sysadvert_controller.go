package controller

import (
	"fish.com/fish-common/entity"
	"fish.com/fish-front-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SysAdvertController struct{}
// @Summary 获取广告
// @Description 获取广告
// @Accept  json
// @Produce json
// @Success 1 {object} entity.Response	"success"
// @Router /sysadvert [get] [post]
func(this *SysAdvertController) List(c *gin.Context){
	if location,err:=strconv.Atoi(c.DefaultQuery("location","0"));err!=nil {
		c.JSON(http.StatusOK,entity.Res.Error())
	}else{
		c.JSON(http.StatusOK,entity.Res.Ok().Result(services.SysAdvertService.Find(int(location))))
	}
}