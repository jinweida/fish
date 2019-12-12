package controller

import (
	"encoding/json"
	"fish.com/fish-common/config"
	"fish.com/fish-common/entity"
	"fish.com/fish-common/util"
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type LoginController struct{}
// @Summary 小程序登录
// @Description 小程序登录
// @Accept  json
// @Produce json
// @Param code query string true "小程序获取code"
// @Success 1 {object} entity.Response	"success"
// @Router /login [get]
func(this *LoginController) Login(c *gin.Context){
	code:=c.DefaultQuery("code","")

	loginUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + url.QueryEscape(config.Conf.WxPay.MiniAppId) +
		"&secret=" + url.QueryEscape(config.Conf.WxPay.MiniSecret) +
		"&js_code=" + url.QueryEscape(code) +
		"&grant_type=authorization_code"

	httpClient := &http.Client{
		Timeout: 15 * time.Second,
	}
	res, err := httpClient.Get(loginUrl)
	if err != nil {
		log4go.Warn("something wrong here: %s", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log4go.Error("httpGet request Body fail")
	}
	//{"session_key":"JoZnwFHJU2mj3YRLC9EaCQ==","openid":"ovm725ebuc6kI0T18hOKhYsCHf0Q"}


	type WechatLogin struct{
		Errcode int;
		Errmsg string;
		Openid string;
		Session_key string ;
		Unionid string;
	}
	wechatLoginResult:=WechatLogin{}
	err=json.Unmarshal(body,&wechatLoginResult)
	fmt.Println(wechatLoginResult)
	if err==nil{
		if wechatLoginResult.Errcode==0{
			fmt.Println("body="+string(body))
			token,_:=util.GenerateToken(wechatLoginResult.Openid,wechatLoginResult.Session_key)
			c.JSON(http.StatusOK,entity.Res.Ok().Result(token))
		}else{
			c.JSON(http.StatusOK,entity.Res.Error().SetMsg(string(body)))
		}
	}else{
		c.JSON(http.StatusOK,entity.Res.Error().SetMsg(string(body)))
	}

}
