package controller

import (
	"fish.com/fish-common/entity"
	"fish.com/fish-front-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MerchantClassifyController struct{}
// @Summary 获取分类
// @Description 获取分类
// @ID get-string-by-int
// @Accept  json
// @Produce json
// @Success 1 {object} entity.Response	"success"
// @Router /merchant/classify [get]
func(this *MerchantClassifyController) List(c *gin.Context){
	c.JSON(http.StatusOK,entity.Res.Ok().Result(services.MerchantClassifyService.Find()))
}
