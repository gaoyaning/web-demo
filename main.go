package main

import (
	"github.com/gin-gonic/gin"
	"web-demo/middleware"
	"web-demo/route"
	"web-demo/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"io"
	"time"
	"path/filepath"
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
	// 以JSON格式为输出，代替默认的ASCII格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	SetRollLogByDay("./logs", "app.log.", logrus.SetOutput)
	// 设置日志等级
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info(config.C)
}

func SetRollLogByDay(dir, prefix string, setOutput func(w io.Writer)) {
	pattern := "2006-01-02"
	today := time.Now().Format(pattern)
	before := today
	fileName := prefix + today
	absDir, _ := filepath.Abs(dir)

	// check log dir existence
	if info, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		logrus.Fatalf("log directory '%s' not exist", absDir)
	} else if err != nil {
		logrus.Fatalf("os.Stat(%s) failed. error: %s", dir, err)
	} else if !info.IsDir() {
		logrus.Fatalf("%s is not a directory", absDir)
	}

	f, err := os.OpenFile(dir+"/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("open log file %s failed. %s", fileName, err)
	}
	setOutput(f)
	go func() {
		oldFile := f
		t := time.NewTicker(time.Second)
		for range t.C {
			now := time.Now().Format(pattern)
			// next day, change log file
			if now != before {
				before = now
				fileName := prefix + now

				newFile, err := os.OpenFile(dir+"/"+fileName, os.O_RDWR|os.O_CREATE, 0666)
				if err != nil {
					// use old log file
					logrus.Errorf("open log file %s failed. %s", fileName, err)
					continue
				}
				// change to new file
				setOutput(newFile)
				// close old file
				if err := oldFile.Close(); err != nil {
					logrus.Errorf("close log file failed. %s", err)
				}
				// set old file to new
				oldFile = newFile
			}
		}
	}()
}
