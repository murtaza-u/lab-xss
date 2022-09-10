package post

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/db"
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
