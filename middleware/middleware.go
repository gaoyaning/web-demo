package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func SetMiddleWare(engine * gin.Engine) {
	engine.Use(logFilter())
}

func  logFilter() gin.HandlerFunc {
	return func(c * gin.Context) {
		start := time.Now().UnixNano()

		requests := make(map[string]string)
		c.Request.ParseForm()
		reqBodyData := c.Request.PostForm
		for key, value := range reqBodyData {
			requests[key] = value[0]
		}

		multiData, err := c.MultipartForm()
		if nil == err {
			data := multiData.Value
			for k, v := range data {
				requests[k] = v[0]
			}
		}
		c.Next()
		ended := time.Now().UnixNano()
		cost := (ended - start)/10e6

		logrus.Infof("request:%v, response:%v, cost_time:%d ms", requests, c.Keys, cost)
	}
}
