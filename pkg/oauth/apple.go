package oauth

import (
	"fmt"
	"github.com/GianOrtiz/apple-auth-go"
)

type AppleUser struct {
	Email string `json:"email"`
}

func (g Provider) AppleValidateOauthJWT(token string) (email string, err error) {
	var appleUser = new(AppleUser)
	user, err := apple.GetUserInfoFromIDToken(token)
	if err != nil {
		return "", fmt.Errorf("failed getting user info: %w", err)
	}
	appleUser.Email = user.Email

	return appleUser.Email, nil
}
