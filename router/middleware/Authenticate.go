package middleware

import (
	"gin-api/app/common/response"
	"github.com/gin-gonic/gin"
)

func Authenticate() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Abort()
		g := response.Gin{Ctx: c}
		g.Response(422, "被中间件 < Authenticate > 拦截", nil)
		return
	}
}
