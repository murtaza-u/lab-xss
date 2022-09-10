package user

import (
	"github.com/labstack/echo/v4"
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

	grp := e.Group("/user")
	grp.POST("/login", login)
	grp.POST("/register", register)
}
