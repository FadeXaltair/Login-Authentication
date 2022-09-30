package models

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	ConfirmPassword string
}

type Login struct {
	gorm.Model
	Email    string
	Password string
}

type CustomClaim struct {
	Name string
	jwt.StandardClaims
}
