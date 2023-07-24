package database

import(
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"

	models "example/models"
)

type DataBaseConnection struct {
	db *gorm.DB
	err error
}

func SetupDb() (*gorm.DB, error) {
	dns := "host=host.docker.internal user=postgres password=postgres port=6432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns))

	if err== nil {
		if rs := db.Exec("SELECT 'CREATE DATABASE DemoDb' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'DemoDb')"); rs.Error != nil {
			return nil,rs.Error
		}
		db.AutoMigrate(&models.People{})
    }
	return db, err
}
