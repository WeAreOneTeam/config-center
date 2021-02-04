package model

import "config-center/conf-core/profile"

type AddProfileBody struct {
	/* 配置项的 key */
	Key string `json:"key"`

	/* 配置项的 value */
	Value string `json:"value"`

	/* 配置项的环境， 测试、开发、灰度、现网 */
	Env string `json:"env"`

	/* 配置项属于哪个（微服务）服务 */
	Service string `json:"service"`

	/* 配置 描述信息 */
	Description string
}

func (body *AddProfileBody) ToProfile() *profile.Profile {
	return &profile.Profile{
		Key:     body.Key,
		Value:   body.Value,
		Env:     body.Env,
		Service: body.Service,
	}
}
