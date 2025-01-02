package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/RichardHoa/go-gin-api/cmd/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(secret []byte, userID int) (string, error) {

	ExpirationTime := time.Second * time.Duration(config.ENVs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(ExpirationTime).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyandGetClaimJWT(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.ENVs.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Use the built-in VerifyExpiresAt function to check expiration
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, fmt.Errorf("token has expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token or claims")

}

func ExtractBearerToken(r *http.Request) (string, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")

	// Check if the Authorization header is in the correct format
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	// Split the header value into "Bearer <token>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("authorization header format must be Bearer <token>")
	}

	// Return the token
	return parts[1], nil
}

func AuthenticateUserToken(r *http.Request) (int, error) {

	token, err := ExtractBearerToken(r)
	if err != nil {
		return 0, err
	}

	claims, err := VerifyandGetClaimJWT(token)
	if err != nil {
		return 0, err
	}

	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid userID")
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return 0, fmt.Errorf("exp claim is missing or invalid")
	}

	// Convert "exp" to time.Time
	expTime := time.Unix(int64(expFloat), 0)

	// Print the expiration date in a readable format
	fmt.Println("Token expiration time:", expTime.Format("2006-01-02 15:04:05"))

	userID := int(userIDFloat)

	return userID, nil

}
