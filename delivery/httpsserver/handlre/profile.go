package handlre

import (
	"fmt"
	"love-date/delivery/httpsserver/response"
	"love-date/service"
	"net/http"
)

type ProfileHandler struct {
	service service.ProfileService
}

func NewProfileHandler(service service.ProfileService) ProfileHandler {

	return ProfileHandler{
		service,
	}
}

func (p ProfileHandler) CreateNewProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createProfileRequest := &service.CreateProfileRequest{}

		dErr := DecodeJSON(r.Body, createProfileRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		newProfile, cErr := p.service.Create(*createProfileRequest)
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("profile created", newProfile.Profile).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this path | %s | isn`t found", r.Method), http.StatusNotFound).ToJSON(w)
	}
}

func (p ProfileHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProfileRequest := &service.GetProfileRequest{}

		dErr := DecodeJSON(r.Body, getProfileRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		profile, cErr := p.service.GetUserProfile(*getProfileRequest)
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("profile loaded", profile.Profile).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this path | %s | isn`t found", r.Method), http.StatusNotFound).ToJSON(w)
	}
}
