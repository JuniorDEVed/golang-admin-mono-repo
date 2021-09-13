package database

import (
	"go-admin/cmd/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open(os.Getenv("PSCALE_DEV")), &gorm.Config{})


	if err != nil {
		panic("Could not connect to the database")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	// export database
	DB = database

}
