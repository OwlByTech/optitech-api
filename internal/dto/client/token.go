package dto

import "github.com/golang-jwt/jwt/v5"

type ClientToken struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}
