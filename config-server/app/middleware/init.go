package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Init(r *gin.Engine) {
	r.Use(Log)
}

func Log(c *gin.Context) {
	log.Printf("ip: %v url: %v", c.ClientIP(), c.FullPath())
}
