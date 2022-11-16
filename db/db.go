package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(postgres.Open("host=localhost user=soib password=soib dbname=week port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// DB.AutoMigrate(&models.Users{})

	
}
