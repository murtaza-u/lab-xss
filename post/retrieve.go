package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/db"
)

func sortByTime(posts []post) []post {
	n := len(posts)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if posts[j].Date.Sub(posts[j+1].Date) > 0 {
				continue
			}

			temp := posts[j]
			posts[j] = posts[j+1]
			posts[j+1] = temp
		}
	}

	return posts
}

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

	posts = sortByTime(posts)

	return ctx.JSON(http.StatusCreated, resp{
		Data:    posts,
		Success: true,
	})
}
