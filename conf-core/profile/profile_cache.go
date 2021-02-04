package profile

/* 缓存接口，用于给查询提供缓存， 可用 redis、memcached 等实现缓存 */
type ConfSCache interface {
	get(key string) string
	Put(conf Profile)
	Remove(key string)
}
