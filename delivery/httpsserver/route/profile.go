package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/delivery/httpsserver/middleware"
	"love-date/service"
	"net/http"
)

func SetProfileRoute(mux *http.ServeMux, repo service.ProfileServiceRepository) {
	profileService := service.NewProfileService(repo)
	profileHandler := handlre.NewProfileHandler(profileService)

	mux.Handle("/profile/create", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.CreateNewProfile)))
	fmt.Println(http.MethodPost + " /profile/create --> create profile route")

	mux.Handle("/profile/get-one", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.GetUserProfile)))
	fmt.Println(http.MethodGet + " /profile/get-one --> get user profile route")

	mux.Handle("/profile/update", middleware.AuthMiddleware(http.HandlerFunc(profileHandler.UpdateProfile)))
	fmt.Println(http.MethodPut + " /profile/update --> update profile route")
}
