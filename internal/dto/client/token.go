package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type ClientToken struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

type ClientTokenResetPassword struct {
	jwt.RegisteredClaims
	ID  int   `json:"id"`
	Exp int64 `json:"exp"`
}
