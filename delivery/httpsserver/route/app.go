package route

import (
	"fmt"
	"love-date/delivery/httpsserver/handler"
	"love-date/delivery/httpsserver/middleware"
	"net/http"
)

func SetAppRoute(mux *http.ServeMux) {
	appHandler := handler.NewAppHandler()
	mux.Handle("/app/special-day", middleware.AuthMiddleware(http.HandlerFunc(appHandler.GetSpecialDays)))
	fmt.Println(http.MethodGet + " app/special-day --> get all special days route")
}
