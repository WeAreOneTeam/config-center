package provider

import (
	"config-center/config-server/app/config"
	"github.com/samuel/go-zookeeper/zk"
	log "github.com/sirupsen/logrus"
	"time"
)

var sessionTimeout = 50 * time.Millisecond

var ZkClient *zk.Conn

func InitZk() {
	ZkClient, ch, err := zk.Connect(config.GetZkAddress(), sessionTimeout)
	if err != nil {
		log.Errorf("zk connect error, err: %v, address: %v, sessionTimeout: %v", config.GetZkAddress(), sessionTimeout)
		panic(err)
	}

	go func() {
		for {
			select {
			case event := <-ch:
				if event.Type == zk.EventNodeCreated {

				}
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()

	log.Infof("InitZk ok, ZkClient: %v \n", ZkClient)
}
