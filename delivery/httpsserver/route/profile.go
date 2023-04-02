package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handler"
	"love-date/delivery/httpsserver/middleware"
	"love-date/service"
	"net/http"
)

func SetProfileRoute(mux *http.ServeMux, profileService *service.ProfileService) {
	profileHandler := handler.NewProfileHandler(*profileService)

	mux.Handle("/profiles/create", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.CreateNewProfile)))
	fmt.Println(http.MethodPost + " /profiles/create --> create profile route")

	mux.Handle("/profiles/get-one", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.GetUserProfile)))
	fmt.Println(http.MethodGet + " /profiles/get-one --> get user profile route")

	mux.Handle("/profiles/update", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.UpdateProfile)))
	fmt.Println(http.MethodPut + " /profiles/update --> update profile route")
}
