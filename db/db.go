package db

import (
	"sync"

	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" 
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
