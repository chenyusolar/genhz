package redis

import (
	"towin/conf"

	"github.com/hertz-contrib/registry/redis"
)

func Init() {
	r := redis.NewRedisRegistry(conf.GetConf().redis.address)
	if r == nil {
		panic("redis registry init failed")
	}
}
