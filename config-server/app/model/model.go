package model

type Config struct {
	Service string `json:"service" binding:"required"`
	Key     string `json:"key" binding:"required"`
	Value   string `json:"value"`
	Version int32  `json:"version"`
	MTime   int64  `json:"mtime"`
}
