package db

import (
	"context"
	"fmt"
	"student-hub-app/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDb struct {
	ctx    context.Context
	client *mongo.Client
	dbName string
}

func (m MongoDb) db() *mongo.Database {
	return m.client.
		Database(m.dbName)
}

func (m *MongoDb) Create(doc model.Document) (*primitive.ObjectID, model.ApiError) {
	result, err := m.db().Collection(doc.CollectionName()).InsertOne(m.ctx, doc)
	if err != nil {
		return nil, model.InternalServerError{Err: fmt.Errorf("error occured when attempting to insert record: %v, err: %v", doc, err)}
	}

	objectID, err := primitive.ObjectIDFromHex(fmt.Sprint(result.InsertedID))
	if err != nil {
		return nil, model.InternalServerError{Err: fmt.Errorf("error retrieving back object id: %v", err)}
	}

	return &objectID, nil
}

func (m *MongoDb) Find(ctx context.Context, id primitive.ObjectID) {

}

func (m *MongoDb) FindAll() {

}

func (m *MongoDb) Update() {

}

func (m *MongoDb) Delete() {

}
