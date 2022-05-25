package main

import (
	"fmt"
	"net/http"

	"iteration-backend/service"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func main() {
	fmt.Println(service.GenerateToken("123"))
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.GET("/login", login)

	e.Logger.Fatal(e.Start(":1323"))
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "jack" && password == "1234" {
		return c.JSON(http.StatusOK, map[string]string{"username": "username", "password": password})
	}
	return fmt.Errorf("error authorize")
}
