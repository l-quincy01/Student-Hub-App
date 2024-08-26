package repository

import (
	"context"
	"fmt"
	"student-hub-app/db"
	"student-hub-app/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const usersCollection = "users"

type UserManager interface {
	Find(ctx context.Context, id primitive.ObjectID) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Insert(ctx context.Context, user model.User) (*primitive.ObjectID, error)
	Update(ctx context.Context, id primitive.ObjectID, field string, value any) (*primitive.ObjectID, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type UserRepository struct {
	mongoDB *db.MongoDB
}

func NewUserManager(mongoDB *db.MongoDB) UserManager {
	return &UserRepository{
		mongoDB: mongoDB,
	}
}

func (u UserRepository) Find(ctx context.Context, id primitive.ObjectID) (model.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	user := new(model.User)
	err := u.mongoDB.Collection(usersCollection).FindOne(ctx, filter).Decode(user)
	if err != nil {
		return model.User{}, fmt.Errorf("unable to find document for collection: %s with id: %s %v", usersCollection, id, err)
	}

	return *user, nil
}

func (u *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	noFilter := bson.D{}
	cursor, err := u.mongoDB.
		Collection(usersCollection).
		Find(ctx, noFilter)

	if err != nil {
		return nil, fmt.Errorf("unable to find documents for collection: %s %v", usersCollection, err)
	}

	var users []model.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("unable to decode items in collection: %s, err:%v", usersCollection, err)
	}

	return users, nil
}

func (u *UserRepository) Insert(ctx context.Context, user model.User) (*primitive.ObjectID, error) {
	result, err := u.mongoDB.
		Collection(usersCollection).
		InsertOne(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("error occured when attempting to insert record: %v, err: %w", user, err)
	}

	objectID := result.InsertedID.(primitive.ObjectID)

	return &objectID, nil
}

func (u *UserRepository) Update(ctx context.Context, id primitive.ObjectID, field string, value any) (*primitive.ObjectID, error) {
	updateFilter := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   field,
					Value: value,
				},
			},
		},
	}

	result, err := u.mongoDB.Collection(usersCollection).UpdateByID(ctx, id, updateFilter)
	if err != nil {
		return nil, fmt.Errorf("unable to update document with object id: %s %v", id, err)
	}

	objectID := result.UpsertedID.(primitive.ObjectID)

	return &objectID, nil
}

func (u *UserRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	idFilter := bson.D{{
		Key:   "_id",
		Value: id,
	}}

	result, err := u.mongoDB.Collection(usersCollection).DeleteOne(ctx, idFilter)
	if err != nil {
		return fmt.Errorf("unable to find documents for collection: %s %v", usersCollection, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no records where found to be deleted with object id: %v", err)
	}

	return nil
}
