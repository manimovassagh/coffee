package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manimovassagh/coffee/database"
	"github.com/manimovassagh/coffee/types"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key") // Replace with your secret key

// JWT middleware
func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtSecret,
	})
}

// Login handler
func LoginHandler(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var user types.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires after 72 hours
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create token"})
	}

	// Set token in cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 72)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged in successfully"})
}
