package database

import "gorm.io/gorm"

type Database struct {
  Client *gorm.DB
}
