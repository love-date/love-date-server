package service

import (
	"love-date/entity"
	"love-date/pkg/errhandling/errmsg"
	"love-date/pkg/errhandling/richerror"
)

type ProfileServiceRepository interface {
	DoesThisUserProfileExist(userID int) (bool, entity.Profile, error)
	CreateProfile(profile entity.Profile) (entity.Profile, error)
	Update(profileID int, profile entity.Profile) (entity.Profile, error)
}
type ProfileService struct {
	repo ProfileServiceRepository
}

func NewProfileService(repo ProfileServiceRepository) ProfileService {
	return ProfileService{repo}
}

type CreateProfileRequest struct {
	Name                    string `json:"name"`
	BirthdayNotifyActive    bool   `json:"birthday_notify_active"`
	SpecialDaysNotifyActive bool   `json:"special_days_notify_active"`
	AuthenticatedUserID     int
}

type CreateProfileResponse struct {
	Profile entity.Profile
}

func (p ProfileService) Create(req CreateProfileRequest) (CreateProfileResponse, error) {
	const op = "profile-service.Create"

	if len(req.Name) < 2 {

		return CreateProfileResponse{}, richerror.New(op).WithMessage("the name's len must be longer than 1").
			WithKind(richerror.KindBadRequest).WithMeta(map[string]interface{}{
			"name": req.Name,
		})
	}

	profileExist, _, err := p.repo.DoesThisUserProfileExist(req.AuthenticatedUserID)
	if err != nil {

		return CreateProfileResponse{}, richerror.New(op).WithWrapError(err)
	}
	if profileExist {

		return CreateProfileResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgHaseBeenProfile).
			WithKind(richerror.KindForbidden)
	}

	if createdProfile, cErr := p.repo.CreateProfile(entity.Profile{
		UserID:                  req.AuthenticatedUserID,
		Name:                    req.Name,
		BirthdayNotifyActive:    req.BirthdayNotifyActive,
		SpecialDaysNotifyActive: req.SpecialDaysNotifyActive,
	}); cErr != nil {

		return CreateProfileResponse{}, richerror.New(op).WithWrapError(cErr)
	} else {

		return CreateProfileResponse{createdProfile}, nil
	}

}

type UpdateProfileRequest struct {
	Name                    *string `json:"name"`
	BirthdayNotifyActive    *bool   `json:"birthday_notify_active"`
	SpecialDaysNotifyActive *bool   `json:"special_days_notify_active"`
	AuthenticatedUserID     int
}

type UpdateProfileResponse struct {
	Profile entity.Profile
}

func (p ProfileService) Update(req UpdateProfileRequest) (UpdateProfileResponse, error) {
	const op = "profile-service"

	profileExist, profile, err := p.repo.DoesThisUserProfileExist(req.AuthenticatedUserID)
	if err != nil {

		return UpdateProfileResponse{}, richerror.New(op).WithWrapError(err)
	}
	if !profileExist {

		return UpdateProfileResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgProfileNotFound).
			WithKind(richerror.KindNotFound)
	}

	if req.Name != nil {
		if len(*req.Name) < 2 {

			return UpdateProfileResponse{}, richerror.New(op).WithMessage("the name's len must be longer than 1").
				WithKind(richerror.KindBadRequest).WithMeta(map[string]interface{}{
				"name": req.Name,
			})
		}
		profile.Name = *req.Name
	}
	if req.SpecialDaysNotifyActive != nil {
		profile.SpecialDaysNotifyActive = *req.SpecialDaysNotifyActive
	}
	if req.BirthdayNotifyActive != nil {
		profile.BirthdayNotifyActive = *req.BirthdayNotifyActive
	}

	if updatedProfile, uErr := p.repo.Update(profile.ID, entity.Profile{
		Name:                    profile.Name,
		BirthdayNotifyActive:    profile.BirthdayNotifyActive,
		SpecialDaysNotifyActive: profile.SpecialDaysNotifyActive,
	}); uErr != nil {

		return UpdateProfileResponse{}, richerror.New(op).WithWrapError(uErr)
	} else {

		return UpdateProfileResponse{updatedProfile}, nil
	}
}

type GetProfileRequest struct {
	AuthenticatedUserID int
}

type GetProfileResponse struct {
	Profile entity.Profile
}

func (p ProfileService) GetUserProfile(req GetProfileRequest) (GetProfileResponse, error) {
	const op = "profile-service.GetUserProfile"

	profileExist, profile, err := p.repo.DoesThisUserProfileExist(req.AuthenticatedUserID)
	if err != nil {

		return GetProfileResponse{}, richerror.New(op).WithWrapError(err)
	}
	if !profileExist {

		return GetProfileResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgProfileNotFound).
			WithKind(richerror.KindNotFound)

	}

	return GetProfileResponse{profile}, nil
}
