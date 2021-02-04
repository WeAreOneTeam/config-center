package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	zkAddr = []string{
		"122.112.203.92:2181",
		"122.112.203.92:2182",
		"122.112.203.92:2183",
	}
)

var zkConn *zk.Conn

func InitConn() {
	zkConn, _, err := zk.Connect(zkAddr, time.Second * 3)
	if err != nil {
		log.Println("connect zk failed")
		return
	}

	result, err := zkConn.Create("/nodes", []byte("192.168.9.213"), 3, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println("create node failed")
		return
	}

	log.Println(result)

	data, stat, err := zkConn.Get(result)
	if err != nil {
		log.Println("get node tt failed")
		return
	}

	log.Println(data)
	log.Println(stat)
	defer zkConn.Close()
	log.Println("connect to zk success")
}

