package middlewares

import (
	"net/http"
	"satrikoding/backend-api/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// get secret key from env
var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header from request
		tokenString := c.GetHeader("Authorization")

		// if token empty return 401 unauthorized
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort() // stop  further requests
			return
		}

		// remove prefix "Bearer " from token
		// the header usually takes the form: "Bearer <token>"
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// create struct to hold the token claims
		claims := &jwt.RegisteredClaims{}

		// parse token and verify signature with jwtKey
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.token) (interface{}, error) {
			// return the secret key to verify the token
			return jwtKey, nil
		})

		// if the token is invalid or an error occurs during parsing
		if err != nil || !token.invalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort() // stop requests
			return
		}

		// save the "sub" (username) claim into context
		c.Set("username", claims.Subject)

		// continuse to the next handler
		c.Next()
	}
}
