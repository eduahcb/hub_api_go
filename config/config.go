package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	entities []interface{}
	db       *gorm.DB
  rdb      *redis.Client
)

func SetEntities(values []interface{}) {
	entities = values
}

func GetDBClient() *gorm.DB {
	return db
}

func GetRedisClient() *redis.Client {
  return rdb
}

func Init() error {
	err := initEnvs()
	if err != nil {
		return err
	}

	db, err = initPostgres()
	if err != nil {
		return err
	}

  rdb = initRedis()

	return nil
}
