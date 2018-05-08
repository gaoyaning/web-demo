package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type logUtil struct {
	logId string
}

var LogUtil *logUtil

func New(c *gin.Engine)  {
}

func (log *logUtil) Infof(format string, args ...interface{}) {
	format = log.logId + "|" + format
	logrus.Infof(format, args...)
}

func (log *logUtil) Info() {

}

func (log *logUtil) Errorf() {

}

func (log *logUtil) Error() {

}

func (log *logUtil) Warrningf() {

}

func (log *logUtil) Warrning() {

}

func (log *logUtil) Debugf() {

}

func (log *logUtil) Debug() {

}