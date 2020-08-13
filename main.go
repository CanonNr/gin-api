package main

import (
	"gin-api/app/server/Web"
	"gin-api/app/server/WebSocket"
	"gin-api/config/yaml"
)

func WebServer() {
	Web.Run(yaml.Conf().Server.Port)
}

func WebSocketServer() {
	WebSocket.Run("8099")
}

func RpcServer() {
}

func main() {
	go WebSocketServer()
	WebServer()
	//RpcServer()
}
