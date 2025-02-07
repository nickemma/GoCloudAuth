package types

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser is a struct that represents the request body for registering a new user

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User is a struct that represents a user in the database
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

// NewUser is a constructor function that returns a new User
func NewUser(registerUser RegisterUser) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 10)
	if err != nil {
		return User{}, err
	}
	// We return a new User with the hashed password
	return User{
		Username:     registerUser.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

// ValidatePassword is a function that validates a password
func ValidatePassword(hashedPassword, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err == nil
}

// CreateToken is a function that creates a JWT token
func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1).Unix()
	// We create a new token with the user's username and the expiration time
	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil,
	}
	// We sign the token with the HMAC256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	secret := os.Getenv("JWT_SECRET")
	// We sign the token with the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Errorf("signingString error %w", err)
		return tokenString
	}
	// We return the token string
	return tokenString
}
