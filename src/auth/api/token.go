package api

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }
}

func getJWTKey() []byte {
    return []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &jwt.MapClaims{
        "Subject":   email,
        "ExpiresAt": expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(getJWTKey())
}