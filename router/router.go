package router

import (
	"gin-api/app/controller/rabbitmq"
	"gin-api/app/controller/rabbitmq/timer"
	"gin-api/app/controller/test"
	"gin-api/router/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": test.Add(1, 3),
		})
	})

	v1 := router.Group("test", middleware.Filter())
	{
		v1.GET("/get", middleware.Authenticate(), test.AddService)
		v1.GET("/data", test.GetData)
		v1.GET("/redis", test.GetRedisData)
		v1.GET("/curl", test.CurlTest)
		v1.GET("/baba", test.BaBa)
	}

	mq := router.Group("mq")
	{
		mq.GET("/put", rabbitmq.Put)
		//mq.GET("/get", rabbitmq.Get)
		mq.GET("/set", timer.Set)

	}

}
