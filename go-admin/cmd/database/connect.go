package database

import (
	"github.com/juniordeved/go-admin/cmd/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("2y449b0yrm05:pscale_pw_hvqtOSio3iU7koGXi1Q_MGY0tmjc1vdf9K6PdqDcPlc@tcp(r74aavhpfqzo.eu-west-2.psdb.cloud)/golang-admin?tls=true"), &gorm.Config{})

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
