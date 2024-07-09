package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document interface {
	CollectionName() string
}

type DocumentBase struct {
	ID primitive.ObjectID `bson:"_id"`
}
