package router

import (
	"fmt"
	"gin-api/app/controller/rabbitmq"
	"gin-api/app/controller/rabbitmq/timer"
	"gin-api/app/controller/test"
	"gin-api/router/middleware"
	"github.com/gin-gonic/gin"
	"log"
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

	flag := router.Group("flag")
	{
		flag.GET("/get", test.Get)
		flag.GET("/get/:type", test.Get)
	}

	router.GET("/t", func(c *gin.Context) {

		//message1 := make(chan int)
		message2 := make(chan int)

		go func() {
			for i := 0; i < 10; i++ {
				flag := <-message2
				if flag > 5 {
					log.Println("message2 已经 > 5")
				}
				//message1 <- i
			}
		}()

		go func() {
			for i := 0; i < 10; i++ {
				message2 <- i
			}
		}()

		for {
			//msg1 := <- message1
			msg2 := <-message2
			//log.Printf("message1: %s \n", string(msg1))
			fmt.Print(msg2)
		}
	})

}
