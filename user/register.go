package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/db"
)

const userBuck = "user"

func register(ctx echo.Context) error {
	uname := ctx.FormValue("username")
	passwd := ctx.FormValue("password")

	if uname == "" || passwd == "" {
		return ctx.JSON(http.StatusBadRequest, resp{
			Err: "incomplete form data",
		})
	}

	h, err := hash([]byte(passwd))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}

	db, err := db.Init(userBuck, dbFile)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}
	defer db.Conn.Close()

	err = db.Put(uname, h)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, resp{
		Success: true, Data: "user created",
	})
}
