package controllers

import (
"net/http"
"santrikoding/backend-api/database"
"santrikoding/backend-api/models"
"santrikoding/backend-api/structs"
"santrikoding/backend-api/helpers"

"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
// Initialize slice to hold user data
var users []models.User

// Get user data from database
database.DB.Find(&users)

// Secd a success response with user data
c.JSON(http.StatusOK, structs.SuccessResponse{
Success: true,
Message: "List Data Users",
Data: users, 
})
}

func CreateUser(c *gin.Context) {
// struct user request
var req = structs.UserCreateRequest{}

// Bind JSON request ke struct UserRequest
if err := c.ShouldBindJSON(&req); err != nil {
c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
Success: false,
Message: "Validation Errors",
Errors: helpers.TranslateErrorMessage(err),
})
return
}

// Initialize new user
user := models.User{
Name: req.Name,
Username: req.Username,
Email: req.Email,
Password: helpers.HashPassword(req.Password),
}

// Save user to database
if err := database.DB.Create(&user).Error; err !=  nil {
c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
Success: false,
Message: "Failed to create user",
Errors: helpers.TranslateErrorMessage(err),
})
return
}

// Send success response
c.JSON(http.StatusCreated, structs.SuccessResponse{
Success: true,
Message: "User created successfuully",
Data: structs.UserResponse{
Id: user.Id,
Name: user.Name,
Username: user.Username,
Email: user.Email,
CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
},
})
}
