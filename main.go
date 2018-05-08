package main

import (
	"github.com/gin-gonic/gin"
	"web-demo/middleware"
	"web-demo/route"
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
	r.Run() // listen and serve on 0.0.0.0:8080
}
