package main

import (
	"santrikoding/backend-api/config"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/routes"
)

func main() {

	// Load config .env
	config.LoadEnv()

	// database initialization
	database.InitDB()

	// setup router
	r := routes.SetupRouter()

	// start server with port 8080
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
