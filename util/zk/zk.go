package zk

import (
	"github.com/docker/libkv/store"
	"github.com/sirupsen/logrus"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store/zookeeper"
	"web-demo/config"
	"strings"
)

var ZK store.Store

func init() {
	servers := config.C.Zk.Servers
	zookeeper.Register()
	if servers == "" {
		logrus.Fatalln("parameter can't be empty")
	}

	serverList := strings.Split(servers, ",")

	var err error
	if ZK, err = libkv.NewStore(store.ZK, serverList, nil); err != nil {
		logrus.Fatalln("connect zookeeper %s failed", serverList, err)
	}
	logrus.Infof("connect zookeeper %s success", serverList)
}