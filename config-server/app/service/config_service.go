package service

import (
	"config-center/config-server/app/model"
)

type IConfigService interface {
	CreateService(service string) error
	DeleteService(service string) error

	CreateConfig(config model.Config) error
	ChangeConfig(config model.Config) error
	DeleteConfig(config model.Config) error

	GetConfig(service string, key string) (model.Config, error)
	GetConfigs(service string) ([]model.Config, error)
}

var ConfigService = ConfigServiceZKImpl{}
