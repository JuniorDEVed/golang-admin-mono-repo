package database

import (
	"github.com/juniordeved/go-admin/cmd/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:012345678910@/go_admin"), &gorm.Config{})

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
