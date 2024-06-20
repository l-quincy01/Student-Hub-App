package model

type Product struct {
	Id          uint64  `json:"id"`
	Price       float32 `json:"price"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}
