package test

import (
	"gin-api/app/common/response"
	"github.com/gin-gonic/gin"
	"log"
)

func Get(c *gin.Context) {
	id := c.Query("id")
	username := c.Query("username")
	// 获取客户端ip
	ipAddr := c.ClientIP()
	// get ua
	ua := c.GetHeader("User-Agent")
	typeName := c.Param("type")

	log.Println(id)
	log.Println(username)
	log.Println(ipAddr)
	log.Println(ua)
	log.Println(typeName)

	g := response.Gin{Ctx: c}
	g.Response(200, "请求成功", nil)
}
