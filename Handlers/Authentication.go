package Handlers

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"../Models"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user Models.User

		c.Bind(&user)

		user, err := Models.GetUserByAuth(user)
		if err != nil {
			return echo.ErrUnauthorized
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})

	}
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.String(http.StatusOK, "Welcome "+name+"!")
	}
}
