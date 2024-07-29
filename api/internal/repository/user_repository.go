package repository

import (
	"fmt"
	"student-hub-app/db"
	"student-hub-app/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Create(user model.User) error
	Read(id primitive.ObjectID) (*model.User, error)
	ReadAll() ([]model.User, error)
	Update(id primitive.ObjectID, field string, value any) error
	Delete(id primitive.ObjectID) error
}

type UserRepositoryImpl struct {
	mongodb db.Mongodb[model.User]
}

const (
	usersCollection = "users"
)

func (u UserRepositoryImpl) Create(user model.User) error {
	userID, err := u.mongodb.Insert(usersCollection, user)
	if err != nil {
		return err
	}

	if userID == nil {
		return fmt.Errorf("unable to create user with object id: %s", user.ID)
	}

	return nil
}

func (u UserRepositoryImpl) Read(id primitive.ObjectID) (*model.User, error) {
	user, err := u.mongodb.Find(usersCollection, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserRepositoryImpl) ReadAll() ([]model.User, error) {
	users, err := u.mongodb.FindAll(usersCollection)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImpl) Update(id primitive.ObjectID, field string, value any) error {
	userID, err := u.mongodb.Update(usersCollection, id, field, value)
	if err != nil {
		return err
	}

	if userID == nil {
		return fmt.Errorf("unable to update user with object id: %s", id)
	}

	return nil
}

func (u UserRepositoryImpl) Delete(id primitive.ObjectID) error {
	err := u.mongodb.Delete(usersCollection, id)
	if err != nil {
		return err
	}

	return nil
}
