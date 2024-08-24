package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB() *MongoDB {
	return new(MongoDB)
}

func (m *MongoDB) Connect(ctx context.Context) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URL")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}
	m.client = client

	return nil
}

func (m *MongoDB) Disconnect(ctx context.Context) error {
	err := m.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) Collection(collection string) *mongo.Collection {
	dbName := os.Getenv("MONGODB_NAME")

	return m.client.Database(dbName).Collection(collection)
}
