package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manimovassagh/coffee/database"
	"github.com/manimovassagh/coffee/types"
	"gorm.io/gorm"
)

// CreateProductHandler handles the creation of new products
func CreateProductHandler(c echo.Context) error {
	userID := c.Get("user_id").(float64) // JWT token's `user_id` is a float64

	var user types.User
	if err := database.DB.Preload("Role").First(&user, uint(userID)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found", "details": err.Error()})
	}

	if user.Role.Name != "seller" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Only sellers can create products"})
	}

	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// type for products
	product := types.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		UserID:      uint(userID),
	}

	if err := database.DB.Create(&product).Error; err != nil {
		if gorm.ErrDuplicatedKey == err {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Product description must be unique for the seller"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create product", "details": err.Error()})
	}

	// Preload the user and role information
	if err := database.DB.Preload("User.Role").First(&product, product.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not load product with user info", "details": err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

// GetProductsBySellerHandler handles retrieving all products created by the logged-in seller
func GetProductsBySellerHandler(c echo.Context) error {
	userID := c.Get("user_id").(float64) // JWT token's `user_id` is a float64

	var products []types.Product
	if err := database.DB.Where("user_id = ?", uint(userID)).Preload("User.Role").Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not retrieve products", "details": err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}
