package connection

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabaseHandler() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	passwd := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DBNAME")
	port := os.Getenv("PG_PORT")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, passwd, dbName, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
