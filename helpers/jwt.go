package helpers

import (
	"santrikoding/backend-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// get scret value from env
var jwtKey = []byte(config.GetEnv("JWT_SECRET",
	"secret_key"))

func GenerateToken(username string) string {
	// Set the token expiration time
	// here we set it to 60 minutes from the current time
	expirationTIme := time.Now().Add(60 * time.Minute)

	// create jwt claims
	// Subject contains the username,
	// and ExpiresAt determines the token expiration time
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// Create a new token with the claims that have been created
	// Using HS256 algrithm to sign token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	// Returns the token in string form
	return token
}
