package handler

import (
	"config-center/config-server/app/consts/code"
	"config-center/config-server/app/model"
	"config-center/config-server/app/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateConfig(c *gin.Context) {
	config := model.Config{}

	err := c.ShouldBind(&config)
	if err != nil {
		log.Errorf("CreateConfig param error, config: %v\n", config)
		FailWithError(c, code.ReqParamInvalid)
	}

	err = service.ConfigService.CreateConfig(config)
	if err != nil {
		log.Errorf("CreateConfig error, config: %v, err: %v\n", config, err)
		FailWithError(c, err)
	}

	Success(c, "")
}

func ChangeConfig(c *gin.Context) {
	config := model.Config{}

	err := c.ShouldBind(&config)
	if err != nil {
		log.Errorf("ChangeConfig param error, config: %v\n", config)
		FailWithError(c, code.ReqParamInvalid)
	}

	err = service.ConfigService.ChangeConfig(config)
	if err != nil {
		log.Errorf("ChangeConfig error, config: %v, err: %v\n", config, err)
		FailWithError(c, err)
	}

	Success(c, "")
}

func DeleteConfig(c *gin.Context) {
	config := model.Config{}

	err := c.ShouldBind(&config)
	if err != nil {
		log.Errorf("DeleteConfig param error, config: %v\n", config)
		FailWithError(c, code.ReqParamInvalid)
	}

	err = service.ConfigService.DeleteConfig(config.Service, config.Key)
	if err != nil {
		log.Errorf("DeleteConfig error, config: %v, err: %v\n", config, err)
		FailWithError(c, err)
	}

	Success(c, "")
}

func GetConfig(c *gin.Context) {
	serviceName := c.Query("service")
	key := c.Query("key")

	if len(serviceName) == 0 || len(key) == 0 {
		log.Errorf("GetConfig param invalid, serviceName: %v, key: %v", serviceName, key)
		FailWithError(c, code.ReqParamInvalid)
	}

	config, err := service.ConfigService.GetConfig(serviceName, key)
	if err != nil {
		log.Errorf("GetConfig error, serviceName: %v , key: %v, err: %v", serviceName, err)
		FailWithError(c, err)
	}

	Success(c, config)
}
