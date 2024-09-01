package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.String(200, "xxx")
	})

	return router

}