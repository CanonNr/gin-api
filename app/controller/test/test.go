package test

import (
	"gin-demo/app/common/response"
	UserModel "gin-demo/app/model/user"
	"gin-demo/app/util/db"
	"github.com/gin-gonic/gin"
)



func AddService(c *gin.Context) {
	result := Add(19,21)
	g := response.Gin{Ctx: c}
	g.Response(200,"请求成功",result)
}

func GetData(c *gin.Context) {
	var user []UserModel.User
	first := db.Db.Find(&user)
	g := response.Gin{Ctx: c}
	g.Response(200,"请求成功",first)
}

func Add(a,b int) int  {
	return a+b
}