package jwtauth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var secret_key string = os.Getenv("SECRET_KEY")

func CreateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(secret_key))

	if err != nil {
		return "", err
	}
	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Unnexpected signing method : %v", token.Header["alg"])

		}
		return []byte(secret_key), nil

	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenstring string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenstring)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid Token")
}
