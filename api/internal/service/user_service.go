package service

import (
	"student-hub-app/internal/model"
	"student-hub-app/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, model.ApiError)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func (u UserServiceImpl) GetAllUsers() ([]model.User, model.ApiError) {
	users, err := u.repo.ReadAll()
	if err != nil {
		return nil, model.InternalServerError{Err: err}
	}

	return users, nil
}
