package jwttoken

import (
	"github.com/golang-jwt/jwt/v4"
	"love-date/config"
	"love-date/pkg/errhandling/errmsg"
	"love-date/pkg/errhandling/richerror"
	"time"
)

var conf *config.Config

type Claims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, email string) (string, error) {
	const op = "jwt.GenerateToken"

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

		return "", richerror.New(op).WithWrapError(err).WithKind(richerror.KindUnexpected).WithMessage(err.Error())
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (bool, Claims, error) {
	const op = "jwt.ValidateJWT"

	conf = config.New()

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(conf.Jwt.Key), nil
	})
	if err != nil {

		if err == jwt.ErrSignatureInvalid {
			return false, Claims{}, richerror.New(op).WithWrapError(err).
				WithMessage(err.Error()).WithKind(richerror.KindBadRequest)
		}

		return false, Claims{}, richerror.New(op).WithWrapError(err).
			WithMessage(err.Error()).WithKind(richerror.KindUnauthorized)

	}

	if !tkn.Valid {

		return false, Claims{}, richerror.New(op).WithMessage(errmsg.ErrorMsgTokenNotValid).WithKind(richerror.KindBadRequest)
	}

	return true, *claims, nil
}
