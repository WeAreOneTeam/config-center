package main

import (
	"config-center/conf-core/profile"
	"fmt"
	//"config-center/conf-management/zookeeper"
	"github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	//zookeeper.InitConn()

	dsn := "conf:conf@tcp(122.112.203.92:5002)/conf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer sqlDB.Close()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	profile := &profile.Profile{
		Id:        uuid.NewV4().String(),
		Key:       "endpoint",
		Value:     "127.0.0.1:1289",
		Env:       "test",
		Service:   "demo",
		CreatedAt: time.Now(),
	}

	result := db.Create(profile)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

}
