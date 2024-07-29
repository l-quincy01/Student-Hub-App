package model

type Author struct {
	Document
	FullName  string     `bson:"full_name" json:"full_name"`
	TextBooks []TextBook `bson:"text_books" json:"text_books"`
}
