package service

import (
	"context"
	"student-hub-app/internal/errors"
	"student-hub-app/internal/model"
	"student-hub-app/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGetter interface {
	GetUsers(ctx context.Context) ([]model.User, errors.ApiError)
}

type UserCreator interface {
	CreateUser(ctx context.Context, user model.User) errors.ApiError
}

type UserUpdater interface {
	UpdateUser(ctx context.Context, id primitive.ObjectID, values map[string]any) errors.ApiError
}

type UserService struct {
	repo repository.UserManager
}

func NewUserCreator(repo repository.UserManager) UserCreator {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUsers(ctx context.Context) ([]model.User, errors.ApiError) {
	users, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, errors.InternalServerError{Err: err}
	}

	return users, nil
}

func (u *UserService) CreateUser(ctx context.Context, user model.User) errors.ApiError {
	_, err := u.repo.Insert(ctx, user)
	if err != nil {
		return errors.InternalServerError{Err: err}
	}

	return nil
}

func (u *UserService) UpdateUser(ctx context.Context, id primitive.ObjectID, values map[string]any) errors.ApiError {
	for field, value := range values {
		_, err := u.repo.Update(ctx, id, field, value)
		if err != nil {
			return errors.InternalServerError{Err: err}
		}
	}

	return nil
}
