package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
)

type Redis struct {
	conn *redis.Client
}

var redisDbLock *Redis
var once sync.Once

func GetInstance() *Redis {

	if redisDbLock == nil {
		once.Do(
			func() {
				c := config.GetInstance()

				r := redis.NewClient(&redis.Options{
					Addr:     c.Redis.HOST,
					Password: c.Redis.PWD,
					DB:       c.Redis.DB,
				})

				redis := &Redis{
					conn: r,
				}

				redisDbLock = redis
			},
		)
	}

	return redisDbLock
}

func (r *Redis) GetDbLock() string {
	return ""
}

func (r *Redis) LockDb() {

}

func (r *Redis) UnlockDb() {

}
