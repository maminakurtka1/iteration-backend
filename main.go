package main

import (
	"iteration-backend/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Login route
	// e.GET("/login", login)
	e.Logger.Fatal(e.Start(":1323"))
}
