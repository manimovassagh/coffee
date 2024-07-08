package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manimovassagh/coffee/database"
	"github.com/manimovassagh/coffee/types"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c echo.Context) error {
	var req types.SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Check if role is valid
	if req.Role != "buyer" && req.Role != "seller" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid role"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not hash password"})
	}

	// Create the user
	user := types.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword), // Store the hashed password
		Role:     types.Role{Name: req.Role},
	}

	// Save the user to the database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	var user types.User

	if err := database.DB.Preload("Role").First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}
