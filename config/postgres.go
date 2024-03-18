package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(Envs.DbURL), &gorm.Config{})
	if err != nil {
    return nil, fmt.Errorf("postgres opening connection error: %v", err)
	}

	err = db.AutoMigrate(entities...)

  if err != nil {
    return nil, fmt.Errorf("postgres migration error: %v", err)
  }

	return db, nil
}
