package handler

import (
	"fmt"
	"love-date/delivery/httpsserver/response"
	"love-date/service"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {

	return UserHandler{
		service,
	}
}

func (p UserHandler) AppendNames(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		userID := r.Context().Value("user_id").(int)

		appendedNames, cErr := p.service.AppendNames(service.AppendPartnerNameRequest{AuthenticatedUserID: userID})
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("", appendedNames.AppendNames).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}
}
