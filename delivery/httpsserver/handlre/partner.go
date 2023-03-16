package handlre

import (
	"fmt"
	"love-date/delivery/httpsserver/response"
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

		newPartner, cErr := p.service.Create(*createPartnerRequest)
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("partner created", newPartner.Partner).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this path | %s | isn`t found", r.Method), http.StatusNotFound).ToJSON(w)
	}
}

func (p PartnerHandler) GetUserPartner(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPartnerRequest := &service.GetUserActivePartnerRequest{}

		dErr := DecodeJSON(r.Body, getPartnerRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		partner, cErr := p.service.GetUserActivePartner(*getPartnerRequest)
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("partner loaded", partner).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this path | %s | isn`t found", r.Method), http.StatusNotFound).ToJSON(w)
	}
}
