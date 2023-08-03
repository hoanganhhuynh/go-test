package database

import(
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"
	"os"
	"fmt"

	models "example/models"
)

type DataBaseConnection struct {
	db *gorm.DB
	err error
}

func SetupDb() (*gorm.DB, error) {

	host := os.Getenv("HOST")
	dbName := os.Getenv("DB_NAME")
	userName := os.Getenv("PS_SERNAME")
	password := os.Getenv("PS_PASSWORD")
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, userName, password, dbName)
	// dns := "host=host.docker.internal user=postgres password=postgres port=6432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns))

	if err== nil {
		if rs := db.Exec("SELECT 'CREATE DATABASE DemoDb' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'DemoDb')"); rs.Error != nil {
			return nil,rs.Error
		}
		db.AutoMigrate(&models.People{})
    }
	return db, err
}
