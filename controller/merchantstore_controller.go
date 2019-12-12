package controller

import (
	"fish.com/fish-common/entity"
	"fish.com/fish-front-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MerchantStoreController struct{}
// @Summary 获取商家店铺
// @Description 获取商家店铺
// @ID get-string-by-int
// @Accept  json
// @Produce json
// @Success 1 {object} entity.Response	"success"
// @Router /merchant/store [get]
func(this *MerchantStoreController) List(c *gin.Context){
	//order
	c.DefaultQuery("assist","0")
	c.DefaultQuery("lng","")
	c.DefaultQuery("lat","")

	//filter
	c.DefaultQuery("sale","")
	c.DefaultQuery("classify_id","")
	c.DefaultQuery("name","")
	c.JSON(http.StatusOK,entity.Res.Ok().Result(services.MerchantStoreService.Find()))
}