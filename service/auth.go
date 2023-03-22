package service

import (
	"fmt"
	"love-date/entity"
	"strings"
)

type AuthServiceRepository interface {
	GoogleValidateOauthJWT(tokenString string) (email string, err error)
}
type AuthService struct {
	repo        AuthServiceRepository
	UserService UserService
}

func NewAuthService(repo AuthServiceRepository, userService UserService) AuthService {

	return AuthService{repo, userService}
}

type ValidateTokenRequest struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
type ValidateTokenResponse struct {
	User entity.User
}

func (a AuthService) RegisterOrLogin(req ValidateTokenRequest) (ValidateTokenResponse, error) {
	var userEmail string

	switch strings.ToLower(req.TokenType) {
	case "google":
		var vErr error
		userEmail, vErr = a.repo.GoogleValidateOauthJWT(req.Token)
		if vErr != nil {

			return ValidateTokenResponse{}, fmt.Errorf("unexpected error: %w", vErr)
		}
	default:

		return ValidateTokenResponse{}, fmt.Errorf("this token type is not supported: %s", req.TokenType)
	}

	userExist, user, uErr := a.UserService.repo.DoesThisUserEmailExist(userEmail)
	if uErr != nil {

		return ValidateTokenResponse{}, fmt.Errorf("unexpected error: %w", uErr)
	}

	if !userExist {
		if createUserResponse, err := a.UserService.Create(UserCreateRequest{
			userEmail,
		}); err != nil {

			return ValidateTokenResponse{}, fmt.Errorf("unexpected error: %w", err)
		} else {
			user = createUserResponse.User
		}
	}

	return ValidateTokenResponse{user}, nil
}
