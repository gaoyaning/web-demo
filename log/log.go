package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
	"fmt"
	"web-demo/config"
	"log"
	"io"
)

const logTimePattern = "2006-01-02 15:04:06.000"
const logDayPattern  = "2006-01-02"

type formatter struct {}
type errorLogHook struct{}
var hookFormatter = formatter{}
var errorLogger = log.Logger{}

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
	s := fmt.Sprintf("%s|%s|%s\n", logTime, logLevel, entry.Message)
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
	switch config.C.Log.Level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
	setDailyRollLog(absPath, "app.log", logrus.SetOutput)

	logrus.AddHook(&errorLogHook{})
	setDailyRollLog(absPath, "app-error.log", errorLogger.SetOutput)
}

func setDailyRollLog(path, fileName string, setOutput func(w io.Writer)) {
	dayTime := time.Now().Format(logDayPattern)
	logFile := fileName + "-" + dayTime
	f, err := os.OpenFile(path+"/"+logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("open log file %s failed. %s", logFile, err)
	}
	setOutput(f)
	go func() {
		oldFile := f
		t := time.NewTicker(time.Second)
		for range t.C {
			nowDay := time.Now().Format(logDayPattern)
			if dayTime != nowDay {
				logFile = fileName + "-" + dayTime
				newFile, err := os.OpenFile(path+"/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if err != err {
					logrus.Fatalf("open log file %s failed. %s", logFile, err)
				}
				setOutput(newFile)
				if err = oldFile.Close(); nil != err {
					logrus.Errorf("close log file failed. %s", err)
				}
				oldFile = newFile
			}
		}
	}()
}

func (hook *errorLogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (hook *errorLogHook) Fire(entry *logrus.Entry) error {
	if content, err := hookFormatter.Format(entry); err == nil {
		errorLogger.Print(string(content))
	} else {
		return err
	}
	return nil
}