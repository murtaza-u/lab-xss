package user

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var secret, dbFile string

type resp struct {
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"err,omitempty"`
	Success bool        `json:"success"`
}

func InitUser(e *echo.Echo, env map[string]string) {
	secret = env["JWT_SECRET"]
	dbFile = env["DB_FILE"]

	e.POST("/login", login)
	e.POST("/register", register)

	cfg := middleware.JWTConfig{
		Claims:     &claims{},
		SigningKey: []byte(secret),
	}

	grp := e.Group("/post")
	grp.Use(middleware.JWTWithConfig(cfg))
}
