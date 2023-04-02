package oauth

import (
	"context"
	"fmt"
	"google.golang.org/api/idtoken"
)

type GoogleUser struct {
	Email string `json:"email"`
}

type Provider struct{}

func NewOauthProvider() Provider {
	return Provider{}
}

func (g Provider) GoogleValidateOauthJWT(token string) (email string, err error) {
	var googleUser = new(GoogleUser)

	payload, err := idtoken.Validate(context.Background(), token, "")
	if err != nil {
		return "", fmt.Errorf("failed getting user info: %w", err)
	}
	googleUser.Email = payload.Claims["email"].(string)

	//response, gErr := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	//if gErr != nil {
	//
	//	return "", fmt.Errorf("failed getting user info: %w", gErr)
	//}
	//contents, err := io.ReadAll(response.Body)
	//if err != nil {
	//	return "", fmt.Errorf("failed reading response body: %s", err.Error())
	//}
	//
	//hasEmail := bytes.Contains(contents, []byte("email"))
	//if !hasEmail {
	//
	//	return "", fmt.Errorf("token hasn't email field")
	//}
	//
	//jErr := json.Unmarshal(contents, googleUser)
	//if jErr != nil {
	//	fmt.Println("err", jErr)
	//}

	return googleUser.Email, nil
}
