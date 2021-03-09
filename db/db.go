package db

import (
	"context"
	"sync"

	"github.com/go-xorm/xorm"
)

var (
	Engine      *xorm.Engine
	once        sync.Once
	RedisClient *Redis
)

func MustInitialize(ctx context.Context, debug bool) {
	once.Do(func() {
		initMySQL()
		initRedis()
	})
}
