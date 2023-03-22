package handlre

import (
	"fmt"
	"love-date/delivery/httpsserver/response"
	"love-date/pkg/jwttoken"
	"love-date/service"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {

	return AuthHandler{authService}
}

func (a AuthHandler) ValidateOauthToken(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		validateTokenRequest := &service.ValidateTokenRequest{}

		dErr := DecodeJSON(r.Body, validateTokenRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		validateTokenResponse, aErr := a.authService.RegisterOrLogin(*validateTokenRequest)
		if aErr != nil {
			response.Fail(aErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
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
