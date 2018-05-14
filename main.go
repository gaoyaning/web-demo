package main

import (
	"github.com/gin-gonic/gin"
	"web-demo/middleware"
	"web-demo/route"
	"web-demo/config"
	"fmt"
	"web-demo/log"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.SetRoute(r)
	middleware.SetMiddleWare(r)
	addr := fmt.Sprintf(":%d", config.C.Port)
	logrus.Info(config.C)
	r.Run(addr)
}

func init() {
	log.InitLog("./logs")
}