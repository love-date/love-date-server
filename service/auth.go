package service

import (
	"fmt"
	"love-date/entity"
)

type AuthServiceRepository interface {
	ValidateOauthJWT(token string) (email string, err error)
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

func (a AuthService) Register(req RegisterRequest) (RegisterResponse, error) {
	userEmail, vErr := a.repo.ValidateOauthJWT(req.tokenString)
	if vErr != nil {

		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", vErr)
	}

	if createUserResponse, err := a.UserService.Create(UserCreateRequest{
		userEmail,
	}); err != nil {

		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	} else {

		return RegisterResponse{createUserResponse.User}, nil
	}
}

type LoginRequest struct {
	tokenString string
}
type LoginResponse struct {
	User entity.User
}

func (a AuthService) login(req LoginRequest) (LoginResponse, error) {
	userEmail, vErr := a.repo.ValidateOauthJWT(req.tokenString)
	if vErr != nil {

		return LoginResponse{}, fmt.Errorf("unexpected error: %w", vErr)
	}

	userExist, user, err := a.UserService.repo.DoesThisUserEmailExist(userEmail)
	if err != nil {

		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	if !userExist {

		return LoginResponse{}, fmt.Errorf("user not found")
	}

	return LoginResponse{user}, nil
}
