package limiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr, password string, db int) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisStorage) Increment(key string, window time.Duration, maxRequests int) (int, error) {
	count, err := r.client.Incr(r.ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if count == 1 {
		r.client.Expire(r.ctx, key, window)
	}

	return int(count), nil
}

func (r *RedisStorage) Block(key string, duration time.Duration) error {
	return r.client.Set(r.ctx, "block:"+key, true, duration).Err()
}

func (r *RedisStorage) IsBlocked(key string) (bool, error) {
	exists, err := r.client.Exists(r.ctx, "block:"+key).Result()
	return exists == 1, err
}
