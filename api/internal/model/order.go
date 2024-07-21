package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Document
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	PurchasedDate time.Time          `bson:"purchased_date" json:"purchased_date"`
}
