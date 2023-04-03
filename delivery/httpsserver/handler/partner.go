package handler

import (
	"fmt"
	"love-date/delivery/httpsserver/response"
	"love-date/errorType"
	"love-date/service"
	"net/http"
)

type PartnerHandler struct {
	service service.PartnerService
}

func NewPartnerHandler(service service.PartnerService) PartnerHandler {

	return PartnerHandler{
		service,
	}
}

func (p PartnerHandler) CreateNewPartner(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		createPartnerRequest := &service.CreatePartnerRequest{}

		dErr := DecodeJSON(r.Body, createPartnerRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		userID := r.Context().Value("user_id").(int)
		createPartnerRequest.AuthenticatedUserID = userID

		newPartner, cErr := p.service.Create(*createPartnerRequest)
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("partner created", newPartner.Partner).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}
}

func (p PartnerHandler) GetUserPartner(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPartnerRequest := &service.GetUserActivePartnerRequest{}

		userID := r.Context().Value("user_id").(int)
		getPartnerRequest.AuthenticatedUserID = userID

		partner, cErr := p.service.GetUserActivePartner(*getPartnerRequest)
		if cErr != nil {
			if cErr == errorType.NotExistData {
				response.Fail(cErr.Error(), http.StatusNoContent).ToJSON(w)

				return
			}
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("partner loaded", partner).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}
}

func (p PartnerHandler) DeleteActivePartner(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		userID := r.Context().Value("user_id").(int)
		success, err := p.service.Remove(service.RemovePartnerRequest{AuthenticatedUserID: userID})
		if err != nil {
			response.Fail(err.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}
		if success {
			response.OK("partner is deleted", nil).ToJSON(w)

			return
		} else {
			response.Fail(err.Error(), http.StatusBadRequest).ToJSON(w)
		}

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}

}

func (p PartnerHandler) UpdatePartner(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		updatePartnerRequest := &service.UpdatePartnerRequest{}
		dErr := DecodeJSON(r.Body, updatePartnerRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		userID := r.Context().Value("user_id").(int)
		updatePartnerRequest.AuthenticatedUserID = userID

		updatedPartner, err := p.service.Update(*updatePartnerRequest)
		if err != nil {
			response.Fail(err.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("partner is updated", updatedPartner.Partner).ToJSON(w)
	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}

}
