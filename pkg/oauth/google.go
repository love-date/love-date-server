package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GoogleUser struct {
	Email string `json:"email"`
}

type OauthProvider struct{}

func NewOauthProvider() OauthProvider {
	return OauthProvider{}
}

func (g OauthProvider) GoogleValidateOauthJWT(token string) (email string, err error) {
	var googleUser = new(GoogleUser)
	response, gErr := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if gErr != nil {

		return "", fmt.Errorf("failed getting user info: %s", gErr.Error())
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body: %s", err.Error())
	}

	hasEmail := bytes.Contains(contents, []byte("email"))
	if !hasEmail {

		return "", fmt.Errorf("token hasn't email field")
	}

	jErr := json.Unmarshal(contents, googleUser)
	if jErr != nil {
		fmt.Println("err", jErr)
	}

	return googleUser.Email, nil
}
