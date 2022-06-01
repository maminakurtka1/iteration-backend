package main

import (
	"iteration-backend/handler"

	_ "iteration-backend/docs"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger Iteration API
// @version 1.0
// @description This is an API for Iteration project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

func setupRouter(db *pgxpool.Pool) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", handler.MainPage)

	// Authorize scope
	e.POST("/sign-up", handler.SignUp(db))
	e.GET("/sign-in", handler.SignIn(db))

	// CV scope
	e.POST("/create-cv", handler.CreateCv(db))

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
