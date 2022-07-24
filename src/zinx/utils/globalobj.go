package utils

import (
	"Zserver/src/zinx/serverinterface"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

/*
	存储全局参数对象
	通过json由用户进行配置
*/

type Config struct {
	TcpServer serverinterface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version        string
	MaxConn        int
	MaxPackageSize uint32 //数据包最大值

	WorkerPoolSize uint32 //消息队列个数
	MaxWorkerTask  uint32 //每个消息队列中最多消息
}

var GlobalConfig *Config

func (g *Config) Reload() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get workspace err: ", err)
	}
	log.Println("Now path is: ", path)
	data, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		log.Fatalln("Read config file err: ", err)
	}
	if err = json.Unmarshal(data, &GlobalConfig); err != nil {
		log.Fatalln("parse config file err: ", err)
	}
}

func init() {
	GlobalConfig = &Config{
		TcpServer:      nil,
		Host:           "0.0.0.0",
		TcpPort:        0,
		Name:           "ZServer",
		Version:        "V0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
		WorkerPoolSize: 10,
		MaxWorkerTask:  1024,
	}
	GlobalConfig.Reload()
}
