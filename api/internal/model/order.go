package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	DocumentBase
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	PurchasedDate time.Time          `bson:"purchased_date" json:"purchased_date"`
}

func (o Order) CollectionName() string {
	return "order_collection"
}
