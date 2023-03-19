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

		userID := r.Context().Value("user_id").(int)
		createProfileRequest.AuthenticatedUserID = userID

		newProfile, cErr := p.service.Create(*createProfileRequest)
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("profile created", newProfile.Profile).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
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

		userID := r.Context().Value("user_id").(int)
		getProfileRequest.AuthenticatedUserID = userID

		profile, cErr := p.service.GetUserProfile(*getProfileRequest)
		if cErr != nil {
			response.Fail(cErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("profile loaded", profile.Profile).ToJSON(w)

	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}
}

func (p ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		updateProfileRequest := &service.UpdateProfileRequest{}
		dErr := DecodeJSON(r.Body, updateProfileRequest)
		if dErr != nil {
			response.Fail(dErr.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		userID := r.Context().Value("user_id").(int)
		updateProfileRequest.AuthenticatedUserID = userID

		updatedProfile, err := p.service.Update(*updateProfileRequest)
		if err != nil {
			response.Fail(err.Error(), http.StatusBadRequest).ToJSON(w)

			return
		}

		response.OK("partner is updated", updatedProfile.Profile).ToJSON(w)
	default:
		response.Fail(fmt.Sprintf("this method | %s | isn`t found at this path", r.Method), http.StatusNotFound).ToJSON(w)
	}

}
