package router

import (
	"gin-demo/app/controller/test"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {


	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": test.Add(1,3),
		})
	})

	router.GET("/test/get",test.AddService)
	router.GET("/test/data",test.GetData)


}
