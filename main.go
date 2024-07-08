package main

import (
	"log"
	"net/http"
	"time"

	"github.com/manimovassagh/coffee/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var db *gorm.DB
	var err error

	// Retry connecting to the database
	for i := 0; i < 10; i++ {
		dsn := "user:password@tcp(localhost:3306)/coffee_app?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retrying... (%d/10)", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatal("Could not connect to the database. Exiting...")
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&types.Role{}, &types.User{}, &types.Category{}, &types.Product{}, &types.Order{}, &types.OrderItem{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migration completed successfully!")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Coffee App!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
