package model

type TextBook struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ISBN     string `json:"isbn"`
	SKU      string `json:"sku"`
	AuthorId int    `json:"authorId"`
}
