package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"os"
)

var (
	config = new(Config)
	env    = os.Getenv("ENV")
)

func Init() {
	loadConfig()
	log.Infof("loadConfig success, config: %#v \n", config)
}

func GetZkAddress() []string {
	return config.address
}

func GetSessionTimeout() int32 {
	return config.sessionTimeout
}

func loadConfig() {
	if env == "" {
		env = "dev"
	}
	configViper := viper.New()
	configViper.AddConfigPath("conf/")
	configViper.SetConfigName(env + ".yml")
	if err := configViper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("unable to read config: %s \n", err))
	}

	if err := configViper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}
}
