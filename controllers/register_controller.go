package controllers

import (
	"net/http"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/helpers"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/structs"

	"github.com/gin-gonic/gin"
)

// Register handles the new user registration process
func Register(c *gin.Context) {
	// Initialize a struct to capture request data
	var req = structs.UserCreateRequest{}

	// Validate JSON request using bindings from Gin
	if err := c.ShouldBindJSON(&req); err != nil {
		// If validation fails, send an error response.
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Create new user data with hashed password
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	// Save user data to database
	if err := database.DB.Create(&user).Error; err != nil {
		// Check if the error is due to duplicate data(e.g. username/email is already registered)
		if helpers.IsDuplicateEntryError(err) {
			// If duplicate, return response 409 Conflict
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	// If successful, send a success response
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
