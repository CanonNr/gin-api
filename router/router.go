package router

import (
	"encoding/json"
	"gin-api/app/controller/rabbitmq"
	"gin-api/app/controller/rabbitmq/timer"
	"gin-api/app/controller/test"
	"gin-api/router/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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

	ws := router.Group("ws")
	{
		ws.GET("/1/:id/:msg", test.WsPush)
	}

	t := router.Group("t")
	{
		// defer
		t.GET("/1", func(context *gin.Context) {
			defer log.Println("All Ending ...")
			log.Println("---------")
			go func() {
				defer log.Println("Thread A Ending")
				for i := 0; i <= 10; i++ {
					time.Sleep(100 * time.Millisecond * time.Duration(1))
					log.Println("Thread A")
				}
			}()

			{
				defer log.Println("Thread B Ending")
				for i := 0; i <= 10; i++ {
					time.Sleep(100 * time.Millisecond * time.Duration(1))
					log.Println("Thread B")
				}
			}

		})
		// make 和 new
		t.GET("/2", func(context *gin.Context) {
			// make 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel；
			// new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针；
		})
		// json
		t.GET("/3", func(context *gin.Context) {

			type Data struct {
				Id   int32
				Name string
			}

			data := new(Data)
			data.Name = "jerry"
			data.Id = 123
			str, _ := json.Marshal(data)

			log.Println(string(str))

			p := new(Data)
			json.Unmarshal(str, p)
			log.Println(p)
		})
	}

}
