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
	userName := os.Getenv("PS_USERNAME")
	password := os.Getenv("PS_PASSWORD")
	port := os.Getenv("PS_PORT")
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, userName, password, dbName)
	
	db, err := gorm.Open(postgres.Open(dns))

	if err== nil {
		db.AutoMigrate(&models.People{})
    }
	return db, err
}
