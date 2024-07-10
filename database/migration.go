package database

import (
	"log"

	"github.com/manimovassagh/coffee/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	// Replace with your actual connection details
	dsn := "user:password@tcp(127.0.0.1:3306)/coffee_app?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop all tables if they exist (use with caution in production)
	err = DB.Migrator().DropTable(&types.Product{}, &types.User{}, &types.Role{}, &types.Order{}, &types.OrderItem{})
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	// Automatically migrate the schema 
	err = DB.AutoMigrate(&types.User{}, &types.Product{}, &types.Role{}, &types.Order{}, &types.OrderItem{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}
}
