package post

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/db"
	"github.com/murtaza-u/lab-xss/id"
	"github.com/murtaza-u/lab-xss/user"
)

func create(ctx echo.Context) error {
	p := new(post)
	err := ctx.Bind(p)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resp{
			Err: err.Error(),
		})
	}

	tkn := ctx.Get("user").(*jwt.Token)
	claims := tkn.Claims.(*user.Claims)
	p.Username = claims.Username

	p.Date = time.Now()

	db, err := db.Init(postBuck, dbFile)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}
	defer db.Conn.Close()

	id, err := id.NewID()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}

	v, err := p.encode()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}

	err = db.Put(id, v)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, resp{
		Data:    p,
		Success: true,
	})
}
