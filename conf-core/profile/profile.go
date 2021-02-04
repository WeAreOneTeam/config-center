package profile

import "time"

type Profile struct {
	/* 配置项的 id */
	Id string `gorm:"column:id"`

	/* 配置项的 key */
	Key string `gorm:"column:conf_key"`

	/* 配置项的 value */
	Value string `gorm:"column:conf_value"`

	/* 配置项的环境， 测试、开发、灰度、现网 */
	Env string `gorm:"column:env"`

	/* 配置项属于哪个（微服务）服务 */
	Service string `gorm:"column:service"`

	/* 配置项描述 */
	Description string `gorm:"column:description"`

	/* 配置项状态 */
	Status string `gorm:"column:status"`

	/* 配置项的版本 */
	Version int `gorm:"column:version"`

	CreatedAt time.Time `gorm:"column:created_at"`

	ModifiedAt time.Time `gorm:"column:modified_at"`

	DeleteAt time.Time `gorm:"column:delete_at"`

	CreateBy string `gorm:"column:created_by"`

	ModifiedBy string `grom:"column:modified_by"`

	DeleteBy string `gorm:"column:delete_by"`
}

func (Profile) TableName() string {
	return "profile"
}

func (p *Profile) GetId() string {
	return p.Id
}

func (p *Profile) GetKey() string {
	return p.Key
}

func (p *Profile) GetEnv() string {
	return p.Env
}

func (p *Profile) GetService() string {
	return p.Service
}

func (p *Profile) GetStatus() string {
	return p.Status
}

func (p *Profile) SetId(id string) {
	p.Id = id
}

func (p *Profile) SetStatus(status string) {
	p.Status = status
}

func (p *Profile) SetCreateBy(operator string) {
	p.CreateBy = operator
}

func (p *Profile) SetModifiedBy(operator string) {
	p.ModifiedBy = operator
}

func (p *Profile) SetDeleteBy(operator string) {
	p.DeleteBy = operator
}

func (p *Profile) String() string {
	return "{key:" + p.Key + " " + "service:" + p.Service + "env" + p.Env + p.Service + "status" + p.Status + "}"
}
