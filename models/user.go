package models

import ("gorm.io/gorm")

type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}

type UserLogin struct {
	gorm.Model
	Email string
	Password string
}