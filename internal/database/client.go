package database

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Database struct {
  Client *gorm.DB
  Redis  *redis.Client
}

