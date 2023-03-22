package handlre

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"love-date/delivery/httpsserver/response"
	"love-date/pkg/jwttoken"
	"love-date/pkg/oauth"
	"love-date/service"
	"net/http"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) AuthHandler {

	return AuthHandler{userService}
}

func (a AuthHandler) ValidateOauthToken(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var validateTokenResponse service.ValidateTokenResponse
		validateTokenRequest := &service.ValidateTokenRequest{}

		dErr := DecodeJSON(r.Body, validateTokenRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		switch validateTokenRequest.TokenType {
		case "google":
			var vErr error

			validateTokenResponse, vErr = a.validateGoogleToken(*validateTokenRequest)
			if vErr != nil {

				response.Fail(vErr.Error(), http.StatusBadRequest).ToJSON(w)

				return
			}
		default:
			response.Fail(fmt.Sprintf("this token type is not supported: %s", validateTokenRequest.TokenType), http.StatusNotFound).ToJSON(w)
		}

		token, jErr := jwttoken.GenerateJWT(validateTokenResponse.User.ID, validateTokenResponse.User.Email)
		if jErr != nil {

			response.Fail(jErr.Error(), http.StatusUnauthorized).ToJSON(w)

			return
		}

		response.OK("token loaded", token).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}
}

func (a AuthHandler) validateGoogleToken(validateTokenRequest service.ValidateTokenRequest) (service.ValidateTokenResponse, error) {
	conf := &oauth2.Config{
		ClientID:     "399793366330-q375vhvmsok3343t7k2qto0mp2r81nks.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:1988/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}

	googleOauthService := oauth.NewGoogleOauth(conf)
	authService := service.NewAuthService(googleOauthService, a.userService)

	validateTokenResponse, cErr := authService.RegisterOrLogin(validateTokenRequest)
	if cErr != nil {
		return service.ValidateTokenResponse{}, cErr
	}

	return validateTokenResponse, nil

}
