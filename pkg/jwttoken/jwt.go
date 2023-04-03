package jwttoken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"love-date/config"
	"net/http"
	"time"
)

var conf *config.Config

type Claims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, email string) (string, error) {
	conf = config.New()

	expirationTime := time.Now().Add(8760 * time.Hour) //One year
	claims := &Claims{
		Email:  email,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(conf.Jwt.Key))
	if err != nil {

		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (bool, Claims, int, error) {
	conf = config.New()

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(conf.Jwt.Key), nil
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
