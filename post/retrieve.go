package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/db"
)

func retrieve(ctx echo.Context) error {
	db, err := db.Init(postBuck, dbFile)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}
	defer db.Conn.Close()

	var posts []post

	data := db.GetAll()
	for _, v := range data {
		p := new(post)
		err := p.decode(v)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, resp{
				Err: err.Error(),
			})
		}

		posts = append(posts, *p)
	}

	return ctx.JSON(http.StatusCreated, resp{
		Data:    posts,
		Success: true,
	})
}
