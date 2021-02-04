package model

import "config-center/conf-core/profile"

type ProfileDto struct {
	/* 配置项的 id */
	Id string `json:"id"`

	/* 配置项的 key */
	Key string `json:"key"`

	/* 配置项的环境， 测试、开发、灰度、现网 */
	Env string `json:"env"`

	/* 配置项属于哪个（微服务）服务 */
	Service string `json:"service"`

	/* 配置项状态 */
	Status string `json:"status"`

	/* 配置项的版本 */
	Version int `json:"version"`
}

func (dto *ProfileDto) From(p *profile.Profile) *ProfileDto{
	return &ProfileDto{
		Id:      p.GetId(),
		Key:     p.GetKey(),
		Env:     p.GetEnv(),
		Service: p.GetService(),
		Status:  p.GetStatus(),
	}
}
