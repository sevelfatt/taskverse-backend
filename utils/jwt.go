package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateJwtToken(userUUID string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userUUID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAndGetJwtTokenClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	expired, err := IsTokenExpired(tokenString)
	if err != nil {
		return nil, err
	}

	if expired {
		return nil, errors.New("token is expired")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}

func GetJwtTokenClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func IsTokenExpired(tokenString string) (bool, error) {
	claims, err := GetJwtTokenClaims(tokenString)
	if err != nil {
		return false, err
	}
	return claims["exp"].(float64) < float64(time.Now().Unix()), nil
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("authorization token is required")
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	return tokenString, nil
}