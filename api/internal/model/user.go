package model

type User struct {
	Document
	FirstName           string      `bson:"first_name" json:"first_name"`
	LastName            string      `bson:"last_name" json:"last_name"`
	PhoneNumber         PhoneNumber `bson:"phone_number" json:"phone_number"`
	PhysicalAddress     Address     `bson:"physical_address" json:"physical_address"`
	PostalAddress       Address     `bson:"postal_address" json:"postal_address"`
	IsPostalAddressSame bool        `bson:"is_postal_address_same" json:"is_postal_address_same"`
}

type Address struct {
	StreetName   string `bson:"street_name" json:"street_name"`
	AddressLine1 string `bson:"address_line1" json:"address_line1"`
	AddressLine2 string `bson:"address_line2" json:"address_line2"`
	City         string `bson:"city" json:"city"`
	Country      string `bson:"country" json:"country"`
	Province     string `bson:"provice" json:"provice"`
	PostalCode   string `bson:"postal_code" json:"postal_code"`
}

