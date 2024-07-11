package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manimovassagh/coffee/database"
	"github.com/manimovassagh/coffee/handlers"
)

func main() {
	// Connect to the database
	database.Connect()

	// Echo instance ini
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users/register", handlers.SignupHandler)
	e.POST("/users/login", handlers.LoginHandler)
	e.POST("/users/refresh", handlers.RefreshTokenHandler)
	e.GET("/userinfo", handlers.UserInfoHandler, handlers.JWTMiddleware)

	e.POST("/products", handlers.CreateProductHandler, handlers.JWTMiddleware)
	e.GET("/products", handlers.GetProductsBySellerHandler, handlers.JWTMiddleware)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
