package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

// ValidateJWTMiddleware is a middleware that validates the JWT token
func ValidateJWTMiddleware(next func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// return a function that takes in a request and returns a response
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// extract the token from the headers
		tokenString := extractTokenFromHeaders(request.Headers)
		if tokenString == "" {
			return events.APIGatewayProxyResponse{
				Body:       "Missing Auth token",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}
		// parse the token and get the claims
		claims, err := parseToken(tokenString)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       "User Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, err
		}

		// check if the token has expired
		expires := int64(claims["expires"].(float64))
		if time.Now().Unix() > expires {
			return events.APIGatewayProxyResponse{
				Body:       "token expired",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}
		// call the next middleware
		return next(request)
	}

}

// extractTokenFromHeaders extracts the token from the headers
func extractTokenFromHeaders(headers map[string]string) string {
	// check if the Authorization header is present
	authHeader, ok := headers["Authorization"]

	if !ok {
		return ""
	}
	// check if the Authorization header has the Bearer token
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	// return the token
	return splitToken[1]
}

// parseToken parses the token and returns the claims
func parseToken(tokenString string) (jwt.MapClaims, error) {
	// parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	// check if there is an error
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}
	// check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid - unauthorized")
	}
	// get the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token of unrecognized type - unauthorized")
	}
	// return the claims
	return claims, nil
}
