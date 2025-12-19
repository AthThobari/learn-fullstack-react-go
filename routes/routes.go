package routes

import (
	"santrikoding/backend-api/controllers"
	"santrikoding/backend-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin initialization
	router := gin.Default()

	// register route
	router.POST("/api/register", controllers.Register)

	// route login
	router.POST("/api/login", controllers.Login)

	// route users
	router.GET("/api/users", middlewares.AuthMiddleware(),
		controllers.FindUser)

	// route user create
	router.POST("/api/users", middlewares.AuthMiddleware(),
		controllers.CreateUser)

	// route user by id
	router.GET("/api/users/:id", middlewares.AuthMiddleware(),
		controllers.FindUserById)

	// route user update
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(),
		controllers.UpdateUser)

	return router
}
