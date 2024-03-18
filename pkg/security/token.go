package security

import (
	"errors"
	"fmt"
	"time"

	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
)

func getExpirationTimeDuration(expirationTime int64) time.Duration {
	expirationTimeDate := time.Unix(expirationTime, 0)
	currentTime := time.Now()

	secondsDifference := expirationTimeDate.Sub(currentTime).Seconds()

	return time.Duration(secondsDifference) * time.Second
}

func Token(userId uint, expirationTime int64, secretKey string) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["userId"] = userId
	permissions["exp"] = expirationTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString, secretKey string) (uint, time.Duration, error) {
	tokenResult, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {

		if errors.Is(err, jwt.ErrTokenMalformed) {
			return 0, 0, &customErrors.TokenMalformed{}
		}

		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return 0, 0, &customErrors.TokenSignatureInvalid{}
		}

		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, 0, &customErrors.TokenExpired{}
		}

		return 0, 0, err
	}

	claims, _ := tokenResult.Claims.(jwt.MapClaims)

	userId := claims["userId"].(float64)
	expirationTime := claims["exp"].(float64)

	expirationTimeDuration := getExpirationTimeDuration(int64(expirationTime))

	return uint(userId), expirationTimeDuration, err
}
