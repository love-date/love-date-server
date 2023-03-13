package service

import (
	"fmt"
	"love-date/entity"
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
	Name                    string
	BirthdayNotifyActive    bool
	SpecialDaysNotifyActive bool
	AuthenticatedUserID     int
}

type CreateProfileResponse struct {
	Profile entity.Profile
}

func (p ProfileService) Create(req CreateProfileRequest) (CreateProfileResponse, error) {
	profileExist, _, err := p.repo.DoesThisUserProfileExist(req.AuthenticatedUserID)
	if err != nil {

		return CreateProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if profileExist {

		return CreateProfileResponse{}, fmt.Errorf("user has been created profile before")
	}

	if createdProfile, cErr := p.repo.CreateProfile(entity.Profile{
		UserID:                  req.AuthenticatedUserID,
		Name:                    req.Name,
		BirthdayNotifyActive:    req.BirthdayNotifyActive,
		SpecialDaysNotifyActive: req.SpecialDaysNotifyActive,
	}); cErr != nil {

		return CreateProfileResponse{}, fmt.Errorf("enexpected error : %w", cErr)
	} else {

		return CreateProfileResponse{createdProfile}, nil
	}

}

type UpdateProfileRequest struct {
	Name                    string
	BirthdayNotifyActive    bool
	SpecialDaysNotifyActive bool
	AuthenticatedUserID     int
}

type UpdateProfileResponse struct {
	Profile entity.Profile
}

func (p ProfileService) Update(req UpdateProfileRequest) (UpdateProfileResponse, error) {
	profileExist, profile, err := p.repo.DoesThisUserProfileExist(req.AuthenticatedUserID)
	if err != nil {

		return UpdateProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if !profileExist {

		return UpdateProfileResponse{}, fmt.Errorf("the profile not found")
	}

	//if profile.UserID != req.AuthenticatedUserID {
	//
	//	return UpdateProfileResponse{}, fmt.Errorf("this user doesn't have permission to update this profile")
	//}

	if updatedProfile, uErr := p.repo.Update(profile.ID, entity.Profile{
		Name:                    req.Name,
		BirthdayNotifyActive:    req.BirthdayNotifyActive,
		SpecialDaysNotifyActive: req.SpecialDaysNotifyActive,
	}); uErr != nil {

		return UpdateProfileResponse{}, fmt.Errorf("unexpected error : %w", uErr)
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
	profileExist, profile, err := p.repo.DoesThisUserProfileExist(req.AuthenticatedUserID)
	if err != nil {

		return GetProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if !profileExist {

		return GetProfileResponse{}, fmt.Errorf("the profile not found")
	}

	return GetProfileResponse{profile}, nil
}
