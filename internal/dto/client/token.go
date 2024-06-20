package dto

import "github.com/golang-jwt/jwt/v5"

type ClientToken struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
