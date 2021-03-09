package db

import (
	"log"

	"ecnu.space/tmp-loser/conf"
	"github.com/gomodule/redigo/redis"
)

const (
	ONE_MINUTE = 60
	ONE_HOUR   = 60 * ONE_MINUTE
	ONE_DAY    = 24 * ONE_HOUR
)

var (
	poolSize = conf.GetAppConfig
)

// Redis redis.
type Redis struct {
	pool    *redis.Pool
	GetConn func() redis.Conn
}

func initRedis() {
	RedisClient = NewRedis(conf.GetAppConfig().RedisConfig.Addr)
}

func NewRedis(addr string) *Redis {
	pool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", addr)
		if err != nil {
			log.Fatal("NewRedis Err: ", err)
		}
		return c, nil
	}, conf.GetAppConfig().RedisConfig.PoolSize)
	return &Redis{
		pool:    pool,
		GetConn: pool.Get,
	}
}
