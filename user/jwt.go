package user

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
