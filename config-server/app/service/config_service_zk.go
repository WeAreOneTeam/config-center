package service

import (
	"config-center/config-server/app/model"
	"config-center/config-server/app/provider"
	log "github.com/sirupsen/logrus"
)

type ConfigServiceZKImpl struct {
}

var pathPrefix = "/service/"

func (configService ConfigServiceZKImpl) CreateService(service string) error {
	var path = getConfigPath(service, "")

	res, err := provider.ZkClient.Create(path, nil, 0, nil)
	if err != nil {
		log.Errorf("Zk Create err: %v, path: %v \n", err, path)
		return err
	}

	log.Infof("CreateService ok, path: %v, res: %v \n", path, res)

	return nil
}

func (configService ConfigServiceZKImpl) DeleteService(service string) error {
	var path = getConfigPath(service, "")

	err := provider.ZkClient.Delete(path, 0)
	if err != nil {
		log.Errorf("Zk Delete err: %v, path: %v \n", err, path)
		return err
	}

	return nil
}

func (configService ConfigServiceZKImpl) CreateConfig(config model.Config) error {
	var path = getConfigPath(config.Service, config.Key)

	res, err := provider.ZkClient.Create(path, []byte(config.Value), 0, nil)
	if err != nil {
		log.Errorf("Zk Create err: %v, path: %v \n", err, path)
		return err
	}

	log.Infof("CreateConfig ok, path: %v, res: %v \n", path, res)

	return nil
}

func (configService ConfigServiceZKImpl) ChangeConfig(config model.Config) error {
	var path = getConfigPath(config.Service, config.Key)

	res, err := provider.ZkClient.Set(path, []byte(config.Value), 0)
	if err != nil {
		log.Errorf("Zk Set err: %v, path: %v \n", err, path)
		return err
	}

	log.Infof("ChangeConfig ok, path: %v, res: %v \n", path, res)

	return nil
}

func (configService ConfigServiceZKImpl) DeleteConfig(service string, key string) error {
	var path = getConfigPath(service, key)

	err := provider.ZkClient.Delete(path, 0)
	if err != nil {
		log.Errorf("Zk Delete err: %v, path: %v \n", err, path)
		return err
	}

	log.Infof("DeleteConfig ok, path: %v \n", path)

	return nil
}

func (configService ConfigServiceZKImpl) GetConfig(service string, key string) (*model.Config, error) {
	var path = getConfigPath(service, key)

	res, stat, err := provider.ZkClient.Get(path)
	if err != nil {
		log.Errorf("Zk Get err: %v, path: %v \n", err, path)
		return nil, err
	}

	config := model.Config{
		Service: service,
		Key:     key,
		Value:   string(res),
		Version: stat.Version,
		MTime:   stat.Mtime,
	}

	log.Infof("GetConfig ok, config: %v", config)

	return &config, nil
}

func (configService ConfigServiceZKImpl) GetConfigs(service string) ([]model.Config, error) {
	var path = getConfigPath(service, "")

	children, _, err := provider.ZkClient.Children(path)
	if err != nil {
		log.Errorf("Zk Children err: %v, path: %v \n", err, path)
		return nil, err
	}

	configs := make([]model.Config, 0, len(children))
	for _, key := range children {
		path = getConfigPath(service, key)
		res, stat, err := provider.ZkClient.Get(path)
		if err != nil {
			log.Errorf("Zk Get err: %v, path: %v \n", err, path)
			return nil, err
		}

		config := model.Config{
			Service: service,
			Key:     key,
			Value:   string(res),
			Version: stat.Version,
			MTime:   stat.Mtime,
		}

		configs = append(configs, config)
	}

	log.Infof("GetAllConfigs ok, service: %v \n", service)

	return configs, nil
}

func getConfigPath(service string, key string) string {
	path := pathPrefix + service
	if len(key) > 0 {
		path += "/" + key
	}

	return path
}
