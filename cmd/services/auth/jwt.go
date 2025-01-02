package auth

import (
	"time"

	"github.com/RichardHoa/go-gin-api/cmd/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(secret []byte, userID int) (string, error) {

	ExpirationTime := time.Second * time.Duration(config.ENVs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"expiredAt": time.Now().Add(ExpirationTime).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
