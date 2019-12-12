package controller

import (
	"fish.com/fish-common/entity"
	"fish.com/fish-front-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SysConfigController struct{
}
// @Summary 获取配置列表
// @Description 获取配置列表
// @ID get-string-by-int
// @Accept  json
// @Produce json
// @Success 1 {object} entity.Response	"success"
// @Router /sysconfig [get]
func (this *SysConfigController) List(c *gin.Context) {
	c.JSON(http.StatusOK,entity.Res.Ok().Result(services.SysConfigService.Find()))
}