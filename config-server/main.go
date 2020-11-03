package main

import (
	"config-center/config-server/app/config"
	"config-center/config-server/app/handler"
	"config-center/config-server/app/middleware"
	"config-center/config-server/app/provider"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	provider.Init()

	r := gin.Default()
	defer r.Run()

	middleware.Init(r)

	register(r)
}

func register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/config", handler.GetConfig)
	r.POST("/config", handler.CreateConfig)
	r.PUT("/config", handler.ChangeConfig)
	r.DELETE("/config", handler.DeleteConfig)

	r.POST("/service/:service", handler.CreateService)
	r.DELETE("/service/:service", handler.DeleteService)
	r.GET("/service/:service", handler.GetConfigs)

}
