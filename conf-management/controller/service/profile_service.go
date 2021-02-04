package service

import (
	"config-center/conf-core/exception"
	"config-center/conf-core/profile"
	log "github.com/sirupsen/logrus"
)

type ProfileService struct {
	storage profile.ConfStorage
}

func (s *ProfileService) GetProfileById(id string) (*profile.Profile, error) {
	p, err := s.storage.GetById(id)
	if err != nil {
		return nil, err
	}

	if p.Status == "deleted" {
		return nil, exception.PROFILE_NOT_EXIST
	}

	return p, nil
}

func (s *ProfileService) GetProfileByKey(service, env, key string) (*profile.Profile, error) {
	p, err := s.storage.GetByKey(service, env, key)
	if err != nil {
		return nil, err
	}

	if p.Status == "deleted" {
		return nil, exception.PROFILE_NOT_EXIST
	}

	return p, nil
}

func (s *ProfileService) AddProfile(profile profile.Profile, operator string) error {
	p, err := s.storage.GetByKey(profile.GetService(), profile.GetEnv(), profile.GetKey())
	if err != nil {
		return err
	}

	if p != nil {
		if p.GetStatus() != "deleted" {
			return exception.PROFILE_EXIST
		}

		/* 相同的配置项反复添加删除， 第二次添加旧纪录删除 */
		_ = s.storage.Remove(profile.GetId())
	}

	err = s.storage.Add(profile, operator)
	if err != nil {
		log.Error("add profile failed. profile: %+v, operator: %s", profile, operator)
		return err
	}

	return nil
}

func (s *ProfileService) UpdateProfile(profile profile.Profile, operator string) error {
	p, err := s.storage.GetById(profile.GetId())
	if err != nil {
		return err
	}

	if p == nil || p.GetStatus() == "deleted" {
		return exception.PROFILE_NOT_EXIST
	}

	err = s.storage.Update(profile, operator)
	if err != nil {
		log.Error("update profile failed. profile: %+v, operator: %s", profile, operator)
		return err
	}

	return nil
}

func (s *ProfileService) deleteProfile(id string, operator string) error {
	p, err := s.storage.GetById(id)
	if err != nil {
		return err
	}

	if p == nil {
		return exception.PROFILE_NOT_EXIST
	}

	err = s.storage.Delete(id, operator)
	if err != nil {
		log.Error("delete profile failed. id: %s, operator: %s", id, operator)
		return err
	}

	return nil
}
