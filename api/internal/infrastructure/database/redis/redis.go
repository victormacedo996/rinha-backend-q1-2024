package redis

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
)

type Redis struct {
	conn *redis.Client
}

var redisDbLock *Redis
var once sync.Once

const DB_LOCK_KEY_NAME = "lock"

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

func (r *Redis) GetDbLock(ctx context.Context) (string, error) {
	val, err := r.conn.Get(ctx, DB_LOCK_KEY_NAME).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *Redis) LockDb(ctx context.Context) error {
	err := r.conn.Set(ctx, DB_LOCK_KEY_NAME, true, 0).Err()
	if err != nil {
		return err
	}

	return nil

}

func (r *Redis) UnlockDb(ctx context.Context) error {
	err := r.conn.Set(ctx, DB_LOCK_KEY_NAME, false, 0).Err()
	if err != nil {
		return err
	}

	return nil

}
