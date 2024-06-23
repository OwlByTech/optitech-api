package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type ClientToken struct {
	ID int32 `json:"id"`
	jwt.RegisteredClaims
}

type ClientTokenResetPassword struct {
	jwt.RegisteredClaims
	ID  int32 `json:"id"`
	Exp int64 `json:"exp"`
}
