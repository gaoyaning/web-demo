package main

import (
	"github.com/gin-gonic/gin"
	"web-demo/middleware"
	"web-demo/route"
	"web-demo/config"
	"fmt"
	"web-demo/util/log"
	_ "web-demo/util/zk/zkMysql"
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
	r.Run(addr)
}

func init() {
	log.InitLog("./logs")
}