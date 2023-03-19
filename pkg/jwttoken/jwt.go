package jwttoken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var jwtKey = []byte("GoLinuxCloudKey")

type Claims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, email string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Email:  email,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (bool, Claims, int, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})
	if err != nil {

		if err == jwt.ErrSignatureInvalid {

			return false, Claims{}, http.StatusUnauthorized, err
		}

		return false, Claims{}, http.StatusBadRequest, err
	}

	if !tkn.Valid {

		return false, Claims{}, http.StatusUnauthorized, fmt.Errorf("token is not valid")
	}

	return true, *claims, 0, nil
}
