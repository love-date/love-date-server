package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handler"
	"love-date/service"
	"net/http"
)

func SetAuthRoute(mux *http.ServeMux, authService *service.AuthService) {

	authHandler := handler.NewAuthHandler(*authService)

	mux.Handle("/validate-token", http.HandlerFunc(authHandler.ValidateOauthToken))
	fmt.Println(http.MethodPost + " /validate-token --> validate token and generate jwt route")
}
