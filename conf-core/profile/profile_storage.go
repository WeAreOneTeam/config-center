package profile

/*
后端存储接口，用于存储配置信息
*/
type ConfStorage interface {
	GetById(id string) (*Profile, error)

	GetByKey(service, env, key string) (*Profile, error)

	Add(conf Profile, operator string) error

	Update(conf Profile, operator string) error

	Delete(id string, operator string) error

	Remove(id string) error
}
