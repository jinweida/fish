package main

import (
	"flag"
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	"fish.com/fish-common/config"
	"fish.com/fish-common/models"
	"fish.com/fish-front-api/controller"
	"fish.com/fish-front-api/routers"
	"github.com/swaggo/files"
	_ "fish.com/fish-front-api/docs"
	"github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func init(){
	log4go.LoadConfiguration("log4go.xml")
	flag.Parse()

	err := config.ParseConf("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("please config config.json")
			os.Exit(0)
		}
		log.Panicln(err)
	}

	//初始化微信客户端
	//    appId：应用ID wxbf6390ac0eaeabf5
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	//client = gopay.NewWeChatClient("wxdaa2ab9ef87b5497", mchId, apiKey, false)

	controller.Client=gopay.NewWeChatClient(config.Conf.WxPay.AppId,config.Conf.WxPay.MchId, config.Conf.WxPay.ApiKey, config.Conf.WxPay.IsProd)
}

// @title 快分API
// @version 1.0
// @description 快分API
// @host http://114.115.204.97:10000
// @BasePath /api
func main() {
	fmt.Println("GoPay Version: ", gopay.Version)
	models.OpenDB(config.Conf)
	gin.SetMode(config.Conf.Mode)




	router := routers.InitRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})


	url := ginSwagger.URL("http://localhost:5080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))	//websocket
	router.Run(fmt.Sprintf(":%d", config.Conf.Server.Port))
}