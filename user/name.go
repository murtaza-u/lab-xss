package user

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func name(ctx echo.Context) error {
	tkn := ctx.Get("user").(*jwt.Token)
	claims := tkn.Claims.(*Claims)
	uname := claims.Username

	return ctx.JSON(http.StatusCreated, resp{
		Data:    echo.Map{"username": uname},
		Success: true,
	})
}
