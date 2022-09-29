package initiializers

import (
	"jwtAuthentication/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDb() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	log.Println("connected")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Login{})
}
