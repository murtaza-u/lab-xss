package post

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/murtaza-u/lab-xss/user"
)

var dbFile string

const postBuck = "post"

type resp struct {
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"err,omitempty"`
	Success bool        `json:"success"`
}

type post struct {
	Username string    `json:"username"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
}

func Init(e *echo.Echo, env map[string]string) {
	dbFile = env["DB_FILE"]

	cfg := middleware.JWTConfig{
		Claims:     &user.Claims{},
		SigningKey: []byte(env["JWT_SECRET"]),
	}

	grp := e.Group("/post")
	grp.GET("/getall", retrieve)
	grp.POST("/create", create, middleware.JWTWithConfig(cfg))
}

func (p post) encode() ([]byte, error) {
	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(p)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (p *post) decode(data []byte) error {
	buff := bytes.NewBuffer(data)
	return gob.NewDecoder(buff).Decode(p)
}
