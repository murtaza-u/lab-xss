package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func logout(ctx echo.Context) error {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Success: false,
			Err:     err.Error(),
		})
	}

	cookie.Expires = time.Now()
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, resp{
		Success: true,
		Data:    "user logged out",
	})
}
