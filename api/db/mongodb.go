package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongodb[T any] struct {
	ctx    context.Context
	client *mongo.Client
}

func NewMongodb[T any](ctx context.Context, client *mongo.Client) Mongodb[T] {
	return Mongodb[T]{
		ctx,
		client,
	}
}

func (m Mongodb[T]) database() *mongo.Database {
	name := os.Getenv("DATABASE_NAME")
	return m.client.Database(name)
}

func (m Mongodb[T]) Find(collection string, id primitive.ObjectID) (*T, error) {
	document := new(T)
	filter := bson.D{{Key: "_id", Value: id}}

	err := m.database().Collection(collection).FindOne(m.ctx, filter).Decode(document)
	if err != nil {
		return nil, fmt.Errorf("unable to find document for collection: %s with id: %s %v", collection, id, err)
	}

	return document, nil
}

func (m Mongodb[T]) FindAll(collection string) ([]T, error) {
	cursor, err := m.database().Collection(collection).Find(m.ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("unable to find documents for collection: %s %v", collection, err)
	}

	var documents []T
	err = cursor.All(context.TODO(), &documents)
	if err != nil {
		return nil, fmt.Errorf("unable to decode items in collection: %s, err:%v", collection, err)
	}

	return documents, nil
}

func (m Mongodb[T]) Insert(collection string, document T) (*primitive.ObjectID, error) {
	result, err := m.database().Collection(collection).InsertOne(m.ctx, document)
	if err != nil {
		return nil, fmt.Errorf("error occured when attempting to insert record: %v, err: %w", document, err)
	}

	objectID, err := primitive.ObjectIDFromHex(fmt.Sprint(result.InsertedID))
	if err != nil {
		return nil, fmt.Errorf("error convertint inserted id to object id: %v", err)
	}

	return &objectID, nil
}

func (m Mongodb[T]) Update(collection string, id primitive.ObjectID, field string, value any) (*primitive.ObjectID, error) {
	update := bson.D{
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

	result, err := m.database().Collection(collection).UpdateByID(m.ctx, id, update)
	if err != nil {
		return nil, fmt.Errorf("unable to update document with object id: %s %v", id, err)
	}

	objectID, err := primitive.ObjectIDFromHex(fmt.Sprint(result.UpsertedID))
	if err != nil {
		return nil, fmt.Errorf("error converting upserted id to object id: %v", err)
	}

	return &objectID, nil
}

func (m Mongodb[T]) Delete(collection string, id primitive.ObjectID) error {
	filter := bson.D{{
		Key:   "_id",
		Value: id,
	}}

	result, err := m.database().Collection(collection).DeleteOne(m.ctx, filter)
	if err != nil {
		return fmt.Errorf("unable to find documents for collection: %s %v", collection, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no records where found to be deleted with object id: %v", err)
	}

	return nil
}
