package service

import (
	"love-date/constant"
	"love-date/entity"
	"love-date/pkg/errhandling/errmsg"
	"love-date/pkg/errhandling/richerror"
	"strings"
)

type AuthServiceRepository interface {
	GoogleValidateOauthJWT(tokenString string) (email string, err error)
	AppleValidateOauthJWT(tokenString string) (email string, err error)
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
	const op = "auth-service.RegisterOrLogin"

	var userEmail string

	switch strings.ToLower(req.TokenType) {
	case constant.GoogleOauthType:
		var vErr error
		userEmail, vErr = a.repo.GoogleValidateOauthJWT(req.Token)
		if vErr != nil {

			return ValidateTokenResponse{}, richerror.New(op).WithWrapError(vErr)
		}
	case constant.AppleOauthType:
		var vErr error
		userEmail, vErr = a.repo.AppleValidateOauthJWT(req.Token)
		if vErr != nil {

			return ValidateTokenResponse{}, richerror.New(op).WithWrapError(vErr)
		}
	default:

		return ValidateTokenResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgNotSupportedTokenType).
			WithKind(richerror.KindBadRequest).WithMeta(map[string]interface{}{
			"token-type": req.TokenType,
		})
	}

	userExist, user, uErr := a.UserService.repo.DoesThisUserEmailExist(userEmail)
	if uErr != nil {

		return ValidateTokenResponse{}, richerror.New(op).WithWrapError(uErr)
	}

	if !userExist {
		if createUserResponse, err := a.UserService.Create(UserCreateRequest{
			userEmail,
		}); err != nil {

			return ValidateTokenResponse{}, richerror.New(op).WithWrapError(uErr)
		} else {
			user = createUserResponse.User
		}
	}

	return ValidateTokenResponse{user}, nil
}
