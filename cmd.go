package labxss

import (
	"errors"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/attack"
	"github.com/murtaza-u/lab-xss/ip"
	"github.com/murtaza-u/lab-xss/post"
	"github.com/murtaza-u/lab-xss/user"
	"github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func loadENV() map[string]string {
	env := make(map[string]string, 3)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	attacker := os.Getenv("ATTACKER_PORT")
	if attacker == "" {
		attacker = "5000"
	}

	dbF := os.Getenv("DB_FILE")
	if dbF == "" {
		dbF = "data.db"
	}

	env["PORT"] = port
	env["ATTACKER_PORT"] = attacker
	env["DB_FILE"] = dbF

	return env
}

var Cmd = &Z.Cmd{
	Name:        "lab-xss",
	Description: "Simple XSS lab",
	Summary:     "Simple XSS lab",
	Commands:    []*Z.Cmd{help.Cmd, appCmd, attackCmd},
}

var appCmd = &Z.Cmd{
	Name:        "app",
	Description: "start the vulnerable application",
	Summary:     "start the vulnerable application",
	Commands:    []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		env := loadENV()

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return errors.New("JWT_SECRET environment variable not set")
		}
		env["JWT_SECRET"] = secret

		e := echo.New()
		e.Static("/", "static")
		user.Init(e, env)
		post.Init(e, env)

		ip, _ := ip.GetLocalIP()
		ip += ":" + env["PORT"]
		log.Printf("App started at %s\n", ip)

		err := e.Start(ip)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	},
}

var attackCmd = &Z.Cmd{
	Name:        "attack",
	Description: "start the attacker",
	Summary:     "start the attacker",
	Commands:    []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		env := loadENV()
		ip, _ := ip.GetLocalIP()
		ip += ":" + env["ATTACKER_PORT"]
		log.Printf("Attacker listening on %s\n", ip)

		return attack.Start(ip)
	},
}
