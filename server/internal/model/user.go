package model

type User struct {
	Id                  int         `json:"id"`
	FirstName           string      `json:"firstName"`
	LastName            string      `json:"lastName"`
	PhoneNumber         PhoneNumber `json:"phoneNumber"`
	PhysicalAddress     Address     `json:"physicalAddress"`
	PostalAddress       Address     `json:"postalAddress"`
	IsPostalAddressSame bool        `json:"isPostalAddressSame"`
}

type Address struct {
	StreetName   string `json:"streetName"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Province     string `json:"provice"`
	PostalCode   string `json:"postalCode"`
}

type PhoneNumber struct {
	Code   CountryCode `json:"code"`
	Digits string      `json:"digits"`
}

func (user User) CollectionName() string {
	return "user_collection"
}

func (profile User) IsSouthAfricanPhoneNumber() bool {
	return profile.PhoneNumber.Code == ZA
}
