package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	"github.com/jinweida/fish-common/config"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"strconv"
	"time"
)
type WXPayController struct{

}
func (this *WXPayController) Test(c *gin.Context){
	//初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gopay.GetRandomString(32))
	bm.Set("body", "小程序测试支付")
	out_trade_no:=strconv.FormatInt(time.Now().UnixNano(),10)
	bm.Set("out_trade_no", out_trade_no)
	bm.Set("total_fee", RandInt(1,5)*100)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "http://www.gopay.ink")
	bm.Set("trade_type", gopay.TradeType_Native)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", gopay.SignType_MD5)
	//bm.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	//嵌套json格式数据（例如：H5支付的 scene_info 参数）
	//h5Info := make(map[string]string)
	//h5Info["type"] = "Wap"
	//h5Info["wap_url"] = "http://www.gopay.ink"
	//h5Info["wap_name"] = "H5测试支付"

	//sceneInfo := make(map[string]map[string]string)
	//sceneInfo["h5_info"] = h5Info
	//
	//bm.Set("scene_info", sceneInfo)

	//参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
	//如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
	sign := gopay.GetWeChatParamSign(config.Conf.WxPay.AppId,config.Conf.WxPay.MchId, config.Conf.WxPay.ApiKey, bm)
	//sign, _ := gopay.GetWeChatSanBoxParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
	bm.Set("sign", sign)

	wxRsp, _ := Client.UnifiedOrder(bm)

	qrcode.WriteFile(wxRsp.CodeUrl, qrcode.Medium, 256, "assets/"+out_trade_no+".png")

	c.File("assets/"+out_trade_no+".png")
	//c.JSON(http.StatusOK,entity.Res.Ok().Result(wxRsp))
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}