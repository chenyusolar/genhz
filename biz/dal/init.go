package dal

import (
	"towin/biz/dal/mysql"
	"towin/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
