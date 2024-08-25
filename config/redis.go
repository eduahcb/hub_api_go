package config

import (
	"github.com/redis/go-redis/v9"
)

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Envs.RedisAddr,
		Password: "",
		DB:       0,
	})

	return rdb
}
