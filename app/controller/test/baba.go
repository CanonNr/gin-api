package test

import (
	"gin-api/app/common/response"
	"github.com/gin-gonic/gin"
)

func BaBa(c *gin.Context) {
	result1 := Add(19, 21)
	result2 := BbDd(32, 24)
	result := Add(result1, result2)
	g := response.Gin{Ctx: c}
	g.Response(200, "请求成功", result)
}
