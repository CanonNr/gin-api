package middleware

import (
	"github.com/gin-gonic/gin"
)

func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Filter ...")
		//c.Abort()
		//c.Next()
	}
}
