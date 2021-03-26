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
	// 单例初始化
	once.Do(func() {
		initMySQL()
		initRedis()
	})
}
