package controllers

import (
	"errors"
	"net/http"
	"anammaulana/backend-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Struct untuk validasi input saat post user
type ValidateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Struct untuk pesan error
type MsgError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Fungsi untuk mendapatkan pesan error berdasarkan tag validator
func GetError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	}
	return "Unknown error"
}

// Handler untuk mendapatkan semua user
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	// Return JSON
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List of Users",
		"data":    users,
	})
}

// Handler untuk menyimpan user baru
func StoreUser(c *gin.Context) {
	var input ValidateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]MsgError, len(ve))
			for i, fe := range ve {
				out[i] = MsgError{fe.Field(), GetError(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		} else {
			// Untuk error lain selain validator
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	// Buat user baru
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	models.DB.Create(&user)

	// Return response JSON
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User Created Successfully",
		"data":    user,
	})
}
