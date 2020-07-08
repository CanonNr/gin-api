package main

import (
	"gin-api/config/websocket"
	"gin-api/config/yaml"
	"gin-api/router"
	"github.com/gin-gonic/gin"
)

func WebServer() *gin.Engine {
	r := gin.Default()
	router.SetupRouter(r)
	return r
}

func WebSocketServer() {
	websocket.Run()
}

func main() {
	go WebServer().Run(yaml.Conf().Server.Port)
	WebSocketServer()
}
