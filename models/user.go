package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Type         string
	Block_status bool
}
type Admin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}


