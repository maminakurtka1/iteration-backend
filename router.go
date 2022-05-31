package main

import (
	"iteration-backend/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupRouter() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/sign-up", handler.SignUp)
	e.GET("/sigh-in", handler.SignIn)

	e.Logger.Fatal(e.Start(":1323"))
}
