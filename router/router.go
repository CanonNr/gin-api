package router

import (
	"gin-api/app/controller/test"
	"gin-api/router/middleware/template"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": test.Add(1, 3),
		})
	})

	router.GET("/test/get", template.SetUp(), test.AddService)
	router.GET("/test/data", test.GetData)
	router.GET("/test/redis", test.GetRedisData)
	router.GET("/test/curl", test.CurlTest)
	router.GET("/test/baba", test.BaBa)

}
