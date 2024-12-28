package database

import (
	"github.com/abnerugeda/go-with-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.Aluno{})
}
