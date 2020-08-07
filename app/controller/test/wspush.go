package test

import (
	"bytes"
	"gin-api/app/server/WebSocket"
	"github.com/gin-gonic/gin"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func WsPush(ctx *gin.Context) {

	id := ctx.Param("id")
	msg := ctx.Param("msg")
	data := bytes.TrimSpace(bytes.Replace([]byte(msg), newline, space, -1))
	message := WebSocket.Message{Id: id, Data: data}
	WebSocket.Send(message)

}
