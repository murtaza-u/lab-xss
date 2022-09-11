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

func Init(e *echo.Echo, env map[string]string) {
	secret = env["JWT_SECRET"]
	dbFile = env["DB_FILE"]

	cfg := middleware.JWTConfig{
		Claims:     &Claims{},
		SigningKey: []byte(secret),
	}

	mid := middleware.JWTWithConfig(cfg)

	e.POST("/login", login)
	e.POST("/register", register)
	e.GET("/logout", logout, mid)
	e.GET("/username", name, mid)

}
