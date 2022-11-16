package main

import (
	"work/db"
	"work/models"
)

func init() {
	db.Connect()
}

func main() {
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Admin{})

}
