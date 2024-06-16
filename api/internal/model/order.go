package model

import "time"

type Order struct {
	Id            uint64    `json:"id"`
	UserId        uint64    `json:"user_id"`
	DatePurchased time.Time `json:"date"`
}
