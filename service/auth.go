package service

import (
	"fmt"
	"love-date/entity"
)

type AuthServiceRepository interface {
	ValidateOauthJWT(tokenString string) (email string, err error)
}
type AuthService struct {
	repo        AuthServiceRepository
	UserService UserService
}

type RegisterRequest struct {
	tokenString string
}
type RegisterResponse struct {
	User entity.User
}

func (a AuthService) RegisterOrLogin(req RegisterRequest) (RegisterResponse, error) {
	userEmail, vErr := a.repo.ValidateOauthJWT(req.tokenString)
	if vErr != nil {

		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", vErr)
	}

	userExist, user, uErr := a.UserService.repo.DoesThisUserEmailExist(userEmail)
	if uErr != nil {

		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", vErr)
	}

	if !userExist {
		if createUserResponse, err := a.UserService.Create(UserCreateRequest{
			userEmail,
		}); err != nil {

			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		} else {
			user = createUserResponse.User
		}
	}

	return RegisterResponse{user}, nil
}
