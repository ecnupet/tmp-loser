package db

import (
	"sync"

	"github.com/go-xorm/xorm"
)

var (
	Engine      *xorm.Engine
	once        sync.Once
	RedisClient *Redis
)

func init() {
	once.Do(func() {
		initMySQL()
		initRedis()
	})
}
