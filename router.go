package main

import (
	"iteration-backend/controller"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func setupRouter() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/sign-up", controller.SignUp)
	e.GET("/sigh-in", controller.SignIn)
}
