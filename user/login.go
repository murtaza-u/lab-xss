package user

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/murtaza-u/lab-xss/db"
)

func login(ctx echo.Context) error {
	uname := ctx.FormValue("username")
	passwd := ctx.FormValue("password")

	if uname == "" || passwd == "" {
		return ctx.JSON(http.StatusBadRequest, resp{
			Err: "incomplete form data",
		})
	}

	db, err := db.Init(userBuck, dbFile)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: err.Error(),
		})
	}
	defer db.Conn.Close()

	if !db.Exists(uname) {
		return ctx.JSON(http.StatusBadRequest, resp{
			Err: "invalid credentials",
		})
	}

	hash := db.Get(uname)
	if hash == nil {
		return ctx.JSON(http.StatusInternalServerError, resp{
			Err: "an internal error occured",
		})
	}

	if !isAuthorized([]byte(passwd), hash) {
		return ctx.JSON(http.StatusBadRequest, resp{
			Err: "invalid credentials",
		})
	}

	exp := time.Now().Add(time.Hour * 6)

	claims := &Claims{
		uname,
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	// create tkn with claims
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token and send it as response.
	tknStr, err := tkn.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tknStr
	cookie.Expires = exp
	cookie.Path = "/"
	// cookie.Secure = true
	// cookie.Domain = "10.0.0.11"
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, resp{
		Success: true,
		Data:    echo.Map{"token": tknStr},
	})
}
