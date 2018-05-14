package config

import (
	"os"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type log struct {
	Level 			string `yml:"level"`
	RedisThreshold int32 	`yml:redisThreshold`
	MysqlThreshold int32 	`yml:mysqlThreshold`
	HttpThreshold  int32 	`yml:httpThreshold`
}

type mysql struct {
	Key		string `yml:"key"`
	Cluster	string `yml:"cluster"`
	ConnTimeout  int32 `yml:"connTimeout"`
	ReadTimeout  int32 `yml:"readTimeout"`
	WriteTimeout int32 `yml:"writeTimeout"`
}

type zk struct {
	Servers	string	`yml:servers`
	Mysql	mysql	`yml:"mysql"`
}

var C = struct {
	Port	int32	`yml:"port"`
	Log	 	log		`yml:"log"`
	Zk		zk		`yml:"zk"`
}{}

func init() {
	Init(&C)
}

func Init(config interface{}) {
	file := ""
	// get config file path from command line parameter
	for i := range os.Args {
		if os.Args[i] == "-c" && len(os.Args) > i+1 {
			file = os.Args[i+1]
			break
		}
	}
	if file == "" {
		logrus.Infoln("not specify config file")
		return
	}

	// get config file content
	content, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Fatalf("read config file %s failed, error: %v", file, err)
	} else {
		logrus.Infof("read config file %s success", file)
	}

	// parse config
	if err := yaml.Unmarshal(content, config); err != nil {
		logrus.Fatalf("parse config file %s to custom struct failed, error: %v", file, err)
	}
}