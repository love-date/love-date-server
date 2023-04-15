package oauth

import (
	"context"
	"google.golang.org/api/idtoken"
	"love-date/pkg/errhandling/richerror"
)

type GoogleUser struct {
	Email string `json:"email"`
}

type Provider struct{}

func NewOauthProvider() Provider {
	return Provider{}
}

func (g Provider) GoogleValidateOauthJWT(token string) (email string, err error) {
	const op = "google.GoogleValidateOauthJWT"

	var googleUser = new(GoogleUser)

	payload, err := idtoken.Validate(context.Background(), token, "")
	if err != nil {

		return "", richerror.New(op).WithWrapError(err).WithMessage(err.Error()).
			WithKind(richerror.KindUnexpected)
	}
	googleUser.Email = payload.Claims["email"].(string)

	return googleUser.Email, nil
}
