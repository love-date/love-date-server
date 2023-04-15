package service

import (
	"fmt"
	"love-date/entity"
	"love-date/pkg/errhandling/errmsg"
	"love-date/pkg/errhandling/richerror"
	"regexp"
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
	Email string `json:"email"`
}
type UserCreateResponse struct {
	User entity.User
}

func (u UserService) Create(req UserCreateRequest) (UserCreateResponse, error) {
	const op = "user-service.Create"

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if len(req.Email) == 0 || !emailRegex.Match([]byte(req.Email)) {

		return UserCreateResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidEmailFormat).
			WithKind(richerror.KindBadRequest).WithMeta(map[string]interface{}{
			"email": req.Email,
		})
	}

	userExist, _, err := u.repo.DoesThisUserEmailExist(req.Email)
	if err != nil {

		return UserCreateResponse{}, richerror.New(op).WithWrapError(err)
	}
	if userExist {

		return UserCreateResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgRegisteredBefore).
			WithKind(richerror.KindBadRequest)
	}

	if createdUser, cErr := u.repo.CreateUser(entity.User{
		Email: req.Email,
	}); cErr != nil {

		return UserCreateResponse{}, richerror.New(op).WithWrapError(cErr)
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
	const op = "user-service.AppendNames"

	profileResult, err := u.profileService.GetUserProfile(GetProfileRequest{req.AuthenticatedUserID})
	if err != nil {
		return AppendPartnerNameResponse{}, richerror.New(op).WithWrapError(err)
	}

	partnerResult, err := u.partnerService.GetUserActivePartner(GetUserActivePartnerRequest{req.AuthenticatedUserID})
	if err != nil {

		return AppendPartnerNameResponse{}, richerror.New(op).WithWrapError(err)
	}

	appendName := fmt.Sprintf("%s %s", profileResult.Profile.Name, partnerResult.Partner.Name)

	return AppendPartnerNameResponse{appendName}, nil
}
