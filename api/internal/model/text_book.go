package model

import "time"

type TextBook struct {
	Product
	Title         string    `json:"title"`
	ISBN          string    `json:"isbn"`
	SKU           string    `json:"sku"`
	Edition       string    `json:"edition"`
	DatePublished time.Time `json:"date_published"`
	AuthorId      uint64    `json:"author_id"`
}
