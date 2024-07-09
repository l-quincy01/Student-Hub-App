package model

type Product struct {
	DocumentBase
	Price       float32 `bson:"price" json:"price"`
	Title       string  `bson:"price" json:"title"`
	Description string  `bson:"description" json:"description"`
}

func (p Product) CollectionName() string {
	return "product_collection"
}
