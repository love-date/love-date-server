package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/service"
	"net/http"
)

func SetAuthRoute(mux *http.ServeMux, userService *service.UserService) {
	authHandler := handlre.NewAuthHandler(*userService)

	mux.Handle("/validate-token", http.HandlerFunc(authHandler.ValidateOauthToken))
	fmt.Println(http.MethodPost + " /validate-token --> validate token and generate jwt route")
}
