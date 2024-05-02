package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabaseHandler() (*gorm.DB, error) {
	dsn := "host=postgres_db user=postgres password=streamkeys9283!@ dbname=streamkeys port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
