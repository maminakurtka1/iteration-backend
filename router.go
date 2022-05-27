package main

import (
	"iteration-backend/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupRouter() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/sign-up", controller.SignUp)
	e.GET("/sigh-in", controller.SignIn)

	e.Logger.Fatal(e.Start(":1323"))
}
