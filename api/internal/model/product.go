package model

type Product struct {
	Document
	Price       float32 `bson:"price" json:"price"`
	Title       string  `bson:"price" json:"title"`
	Description string  `bson:"description" json:"description"`
}