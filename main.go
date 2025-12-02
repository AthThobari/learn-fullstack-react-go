package main

import (
	"github.com/gin-gonic/gin"

	"santrikoding/backend-api/config"
	"santrikoding/backend-api/database"
)

func main() {

	// Load config .env
	config.LoadEnv()

	// Gin's initialization
	router := gin.Default()

	// database initialization
	database.InitDB()

	// create a route with the GET method
	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// start server with port 8080
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
