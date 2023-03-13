package service

import (
	"fmt"
	"love-date/entity"
)

type UserServiceRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	DoesThisUserEmailExist(email string) (bool, entity.User, error)
}

type UserService struct {
	repo           UserServiceRepository
	partnerService PartnerService
	profileService ProfileService
}

func NewUserService(repo UserServiceRepository, partnerService PartnerService,
	profileService ProfileService) UserService {
	return UserService{
		repo,
		partnerService,
		profileService,
	}
}

type UserCreateRequest struct {
	Email string
}
type UserCreateResponse struct {
	User entity.User
}

func (u UserService) Create(req UserCreateRequest) (UserCreateResponse, error) {
	userExist, _, err := u.repo.DoesThisUserEmailExist(req.Email)
	if err != nil {

		return UserCreateResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if userExist {

		return UserCreateResponse{}, fmt.Errorf("the email has been registered before")
	}

	if createdUser, cErr := u.repo.CreateUser(entity.User{
		Email: req.Email,
	}); cErr != nil {

		return UserCreateResponse{}, fmt.Errorf("unexpected error: %w", cErr)
	} else {

		return UserCreateResponse{createdUser}, nil
	}

}

type AppendPartnerNameRequest struct {
	AuthenticatedUserID int
}

type AppendPartnerNameResponse struct {
	AppendNames string
}

func (u UserService) AppendNames(req AppendPartnerNameRequest) (AppendPartnerNameResponse, error) {
	profileResult, err := u.profileService.GetUserProfile(GetProfileRequest{req.AuthenticatedUserID})
	if err != nil {

		return AppendPartnerNameResponse{}, fmt.Errorf("can't get profile : %w", err)
	}

	partnerResult, err := u.partnerService.GetUserActivePartner(GetUserActivePartnerRequest{req.AuthenticatedUserID})
	if err != nil {

		return AppendPartnerNameResponse{}, fmt.Errorf("can't get partner : %w", err)
	}

	appendName := fmt.Sprintf("%s %s", profileResult.Profile.Name, partnerResult.Partner.Name)

	return AppendPartnerNameResponse{appendName}, nil
}
