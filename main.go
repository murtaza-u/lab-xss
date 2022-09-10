package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/user"
)

type resp struct {
	Err     error `json:"err,omitempty"`
	Success bool  `json:"success"`
	Data    any   `json:"data,omitempty"`
}

func loadENV() map[string]string {
	env := make(map[string]string, 3)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbF := os.Getenv("DB_FILE")
	if dbF == "" {
		dbF = "data.db"
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	env["PORT"] = port
	env["DB_FILE"] = dbF
	env["JWT_SECRET"] = secret

	return env
}

func main() {
	env := loadENV()

	e := echo.New()
	user.InitUser(e, env)

	err := e.Start(":" + env["PORT"])
	if err != nil {
		log.Fatal(err)
	}
}
