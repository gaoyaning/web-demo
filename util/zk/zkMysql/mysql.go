package zkMysql

import (
	"web-demo/config"
	"web-demo/util/zk"
	"github.com/sirupsen/logrus"
)

func init() {
	zkMysqlKey := config.C.Zk.Mysql.Key
	if isExist, err := zk.ZK.Exists(zkMysqlKey); !isExist || nil != err {
		logrus.Fatalf("can not find mysql node: %s or check failed error: %v", zkMysqlKey, err)
		return
	}

	stopCh := make(<-chan struct{})
	events, err := zk.ZK.Watch(zkMysqlKey, stopCh)
	if nil != err {
		logrus.Fatalf("watch mysql node: %s error: %v", zkMysqlKey, err)
		return
	}
	for {
		select {
		case pair := <-events:
			// Do something with events
			logrus.Infof("value changed on key %v: new value=%v", zkMysqlKey, pair.Value)
		}
	}
}