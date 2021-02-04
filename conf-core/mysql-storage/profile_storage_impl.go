package mysql_storage

import (
	"config-center/conf-core/exception"
	"config-center/conf-core/profile"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfileStorageImpl struct {
	db *gorm.DB
}

func (p *ProfileStorageImpl) GetById(id string) (*profile.Profile, error) {
	if id == "" {
		log.Error("id is null")
		return nil, exception.INTERVAL_ERROR
	}

	var pf profile.Profile
	if err := p.db.Model(&profile.Profile{}).Where(&profile.Profile{Id: id}).Find(pf).Error; err != nil {
		log.Error("query profile failed id: %d, err: %s", id, err)
		return nil, exception.INTERVAL_ERROR
	}

	return &pf, nil
}

func (p *ProfileStorageImpl) GetByKey(service, env, key string) (*profile.Profile, error) {
	if service == "" || env == "" || key == "" {
		log.Error("param invalid, service: %s, env: %s, key: %s", service, env, key)
		return nil, exception.INTERVAL_ERROR
	}

	var pf profile.Profile
	if err := p.db.Model(&profile.Profile{}).Where(&profile.Profile{Service: service, Env: env, Key: key}).Find(&pf).Error; err != nil {
		log.Error("query profile failed service: %s, env: %s, key: %s, err: %v", service, env, key, err)
		return nil, exception.INTERVAL_ERROR
	}

	return &pf, nil
}

func (p *ProfileStorageImpl) Add(conf profile.Profile, operator string) error {
	conf.SetCreateBy(operator)
	conf.SetStatus("enabled")
	if err := p.db.Model(&profile.Profile{}).Create(&conf); err != nil {
		return exception.INTERVAL_ERROR
	}

	return nil
}

func (p *ProfileStorageImpl) Update(conf profile.Profile, operator string) error {
	conf.SetModifiedBy(operator)
	if err := p.db.Model(&profile.Profile{}).Where(&profile.Profile{Id: conf.GetId()}).Updates(conf).Error; err != nil {
		return exception.INTERVAL_ERROR
	}

	return nil
}

func (p *ProfileStorageImpl) Delete(id string, operator string) error {
	if id == "" {
		log.Error("id is null")
		return exception.INTERVAL_ERROR
	}

	if err := p.db.Model(&profile.Profile{}).Where(&profile.Profile{Id: id}).Updates(profile.Profile{Status: "deleted", DeleteBy: operator}).Error; err != nil {
		return exception.INTERVAL_ERROR
	}

	return nil
}

func (p *ProfileStorageImpl) Remove(id string) error {
	if id == "" {
		log.Error("id is null")
		return exception.INTERVAL_ERROR
	}

	if err := p.db.Model(&profile.Profile{}).Where(&profile.Profile{Id: id}).Delete(profile.Profile{}).Error; err != nil {
		return exception.INTERVAL_ERROR
	}

	return nil
}
