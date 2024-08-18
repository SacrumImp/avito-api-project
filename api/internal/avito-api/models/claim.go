package models

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	Role string
	jwt.RegisteredClaims
}
