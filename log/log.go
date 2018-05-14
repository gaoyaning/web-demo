package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
	"fmt"
)

const logTimePattern = "2006-01-02 15:04:06.000"

type formatter struct {}

func (formatter) Format(entry *logrus.Entry) ([]byte, error) {
	logTime  := time.Now().Format(logTimePattern)
	logLevel := "DEBUG"
	switch entry.Level {
	case logrus.PanicLevel:
		logLevel = "PANIC"
	case logrus.FatalLevel:
		logLevel = "FATAL"
	case logrus.ErrorLevel:
		logLevel = "ERROR"
	case logrus.InfoLevel:
		logLevel = "INFO"
	case logrus.WarnLevel:
		logLevel = "WARNING"
	case logrus.DebugLevel:
		logLevel = "DEBUG"
	}
	s := fmt.Sprintf("%s|%s|%s|", logTime, logLevel, entry.Message)
	return []byte(s), nil
}

func InitLog(path string) {
	absPath, _ := filepath.Abs(path)
	if _, err := os.Stat(absPath); nil != err || os.IsNotExist(err) {
		err = os.MkdirAll(absPath, 0774)
		if nil != err {
			logrus.Fatalf("logs is not exist and create path fail, error: %v", err)
			return
		}
	}
	logrus.SetFormatter(formatter{})
}