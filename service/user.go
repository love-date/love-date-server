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
	repo UserServiceRepository
}

func NewUserService(repo UserServiceRepository) UserService {
	return UserService{repo}
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
