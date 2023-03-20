package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/delivery/httpsserver/middleware"
	"love-date/service"
	"net/http"
)

func SetPartnerRoute(mux *http.ServeMux, repo service.PartnerServiceRepository) {
	partnerService := service.NewPartnerService(repo)
	partnerHandler := handlre.NewPartnerHandler(partnerService)

	mux.Handle("/partner/create", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.CreateNewPartner)))
	fmt.Println(http.MethodPost + " /partner/create --> create partner route")

	mux.Handle("/partner/get-active", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.GetUserPartner)))
	fmt.Println(http.MethodGet + " /partner --> get user active partner route")

	mux.Handle("/partner/update", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.UpdatePartner)))
	fmt.Println(http.MethodPut + " /partner/update --> update partner route")

	mux.Handle("/partner/delete", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.DeleteActivePartner)))
	fmt.Println(http.MethodGet + " /partner --> delete user active partner route")

}
