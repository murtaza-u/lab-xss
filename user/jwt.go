package user

import "github.com/golang-jwt/jwt"

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
