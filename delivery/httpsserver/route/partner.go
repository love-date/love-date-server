package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/delivery/httpsserver/middleware"
	"love-date/service"
	"net/http"
)

func SetPartnerRoute(mux *http.ServeMux, partnerService *service.PartnerService) {
	partnerHandler := handlre.NewPartnerHandler(*partnerService)

	mux.Handle("/partners/create", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.CreateNewPartner)))
	fmt.Println(http.MethodPost + " /partners/create --> create partner route")

	mux.Handle("/partners/get-active", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.GetUserPartner)))
	fmt.Println(http.MethodGet + " /partners/get-active --> get user active partner route")

	mux.Handle("/partners/update", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.UpdatePartner)))
	fmt.Println(http.MethodPut + " /partners/update --> update partner route")

	mux.Handle("/partners/delete", middleware.AuthMiddleware(http.HandlerFunc(partnerHandler.DeleteActivePartner)))
	fmt.Println(http.MethodGet + " /partners/delete --> delete user active partner route")
}
