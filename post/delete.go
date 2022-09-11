package post

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/db"
	"github.com/murtaza-u/lab-xss/user"
)

func delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if len(id) != 10 {
		return ctx.JSON(http.StatusBadRequest, resp{
			Err: "invalid ID",
		})
	}

	db, err := db.Init(postBuck, dbFile)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}
	defer db.Conn.Close()

	tkn := ctx.Get("user").(*jwt.Token)
	claims := tkn.Claims.(*user.Claims)

	p := new(post)
	p.decode(db.Get(id))
	if p.Username != claims.Username {
		return ctx.JSON(http.StatusUnauthorized, resp{
			Err: "unauthorized",
		})
	}

	err = db.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, resp{
		Success: true,
		Data:    fmt.Sprintf("Post deleted: %s", id),
	})
}
