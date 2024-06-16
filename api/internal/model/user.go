package model

type User struct {
	Id                  uint64      `json:"id"`
	FirstName           string      `json:"first_name"`
	LastName            string      `json:"last_name"`
	PhoneNumber         PhoneNumber `json:"phone_number"`
	PhysicalAddress     Address     `json:"physical_address"`
	PostalAddress       Address     `json:"postal_address"`
	IsPostalAddressSame bool        `json:"is_postal_address_same"`
}

type Address struct {
	StreetName   string `json:"street_name"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Province     string `json:"provice"`
	PostalCode   string `json:"postal_code"`
}

func (u User) CollectionName() string {
	return "user_collection"
}

func (u User) IsSouthAfricanPhoneNumber() bool {
	return u.PhoneNumber.Code == ZA
}
