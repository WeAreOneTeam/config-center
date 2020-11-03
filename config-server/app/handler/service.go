package handler

import (
	"config-center/config-server/app/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateService(c *gin.Context) {
	serviceName := c.Param("service")

	err := service.ConfigService.CreateService(serviceName)
	if err != nil {
		log.Errorf("CreateService error, serviceName: %v , err: %v", serviceName, err)
		FailWithError(c, err)
	}

	Success(c, "")
}

func DeleteService(c *gin.Context) {
	serviceName := c.Param("service")

	err := service.ConfigService.DeleteService(serviceName)
	if err != nil {
		log.Errorf("DeleteService error, serviceName: %v , err: %v", serviceName, err)
		FailWithError(c, err)
	}

	Success(c, "")
}

func GetConfigs(c *gin.Context) {
	serviceName := c.Param("service")

	configs, err := service.ConfigService.GetConfigs(serviceName)
	if err != nil {
		log.Errorf("DeleteService error, serviceName: %v , err: %v", serviceName, err)
		FailWithError(c, err)
	}

	Success(c, configs)
}
