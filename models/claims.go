package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"userType"`
	jwt.StandardClaims
}
