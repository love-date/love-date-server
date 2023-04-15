package oauth

import (
	"github.com/GianOrtiz/apple-auth-go"
	"love-date/pkg/errhandling/richerror"
)

type AppleUser struct {
	Email string `json:"email"`
}

func (g Provider) AppleValidateOauthJWT(token string) (email string, err error) {
	const op = "apple.AppleValidateOauthJWT"

	var appleUser = new(AppleUser)
	user, err := apple.GetUserInfoFromIDToken(token)
	if err != nil {

		return "", richerror.New(op).WithWrapError(err).WithMessage(err.Error()).
			WithKind(richerror.KindUnexpected)
	}
	appleUser.Email = user.Email

	return appleUser.Email, nil
}
