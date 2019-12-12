package middleware

import (
	"fish.com/fish-common/entity"
	"fish.com/fish-common/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func JWTAuth() gin.HandlerFunc{
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, entity.Res.Expire().SetMsg("token is nil"))
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// parseToken 解析token包含的信息
		claims, err := util.ParseToken(token)
		if err != nil {
			if err == util.TokenExpired {
				c.JSON(http.StatusOK, entity.Res.Expire())
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, entity.Res.Expire())
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}