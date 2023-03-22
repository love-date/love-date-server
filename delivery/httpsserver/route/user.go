package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/delivery/httpsserver/middleware"
	"love-date/service"
	"net/http"
)

func SetUserRoute(mux *http.ServeMux, userService *service.UserService) {
	userHandler := handlre.NewUserHandler(*userService)

	mux.Handle("/users/append-name", middleware.AuthMiddleware(http.HandlerFunc(userHandler.AppendNames)))
	fmt.Println(http.MethodGet + " /users/append-name --> get append user and partner names route")

}
