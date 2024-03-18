package config

import "gorm.io/gorm"

var (
	entities []interface{}
	db       *gorm.DB
)

func SetEntities(values []interface{}) {
	entities = values
}

func GetDBClient() *gorm.DB {
	return db
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

	return nil
}
