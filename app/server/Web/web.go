package Web

import (
	"gin-api/router"
	"github.com/gin-gonic/gin"
)

func Run(port string) *gin.Engine {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(port)
	return r
}
