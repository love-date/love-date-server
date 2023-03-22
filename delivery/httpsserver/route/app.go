package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handlre"
	"love-date/delivery/httpsserver/middleware"
	"net/http"
)

func SetAppRoute(mux *http.ServeMux) {
	appHandler := handlre.NewAppHandler()

	mux.Handle("app/special-day", middleware.AuthMiddleware(http.HandlerFunc(appHandler.GetSpecialDays)))
	fmt.Println(http.MethodPost + " app/special-day --> get all special days route")
}
