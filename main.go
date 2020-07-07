package main

import (
	"gin-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 路由
	router.SetupRouter(r)

	r.Run()
}