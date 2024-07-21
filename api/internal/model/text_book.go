package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TextBook struct {
	Product
	Title         string             `bson:"title" json:"title"`
	ISBN          string             `bson:"isbn" json:"isbn"`
	SKU           string             `bson:"sku" json:"sku"`
	Edition       string             `bson:"edition" json:"edition"`
	DatePublished time.Time          `bson:"date_published" json:"date_published"`
	AuthorID      primitive.ObjectID `bson:"author_id" json:"author_id"`
}
