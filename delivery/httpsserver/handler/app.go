package handler

import (
	"fmt"
	"love-date/delivery/httpsserver/response"
	"love-date/pkg/specialday"
	"net/http"
)

type AppHandler struct {
}

func NewAppHandler() AppHandler {

	return AppHandler{}
}

func (p AppHandler) GetSpecialDays(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		specialDays, sErr := specialday.GetSpecialDays()
		if sErr != nil {
			response.Fail(sErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("special days loaded", specialDays).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}
}
