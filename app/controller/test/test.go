package test

import (
	"gin-api/app/common/response"
	UserModel "gin-api/app/model/user"
	"gin-api/app/util/db"
	"gin-api/app/util/redis"
	"github.com/gin-gonic/gin"
	Redigo "github.com/gomodule/redigo/redis"
)

func AddService(c *gin.Context) {
	result := Add(19, 21)
	g := response.Gin{Ctx: c}
	g.Response(200, "请求成功", result)
}

func GetData(c *gin.Context) {
	var user []UserModel.User
	first := db.Db.Find(&user)
	g := response.Gin{Ctx: c}
	g.Response(200, "请求成功", first)
}

func GetRedisData(c *gin.Context) {
	var result [2]interface{}
	client := redis.Client
	client.Do("SET", "username", "baba")
	var data, err = client.Do("GET", "username")

	result[0], _ = Redigo.String(data, nil)
	result[1] = err
	g := response.Gin{Ctx: c}
	g.Response(200, "请求成功", result)
}
func Add(a, b int) int {
	return a + b
}
