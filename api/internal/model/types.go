package model

import "fmt"

type OAuthUserInfo struct {
	Id             string `json:"id"`
	FullName       string `json:"name"`
	FirstName      string `json:"given_name"`
	LastName       string `json:"family_name"`
	Email          string `json:"email"`
	VerifiedEmail  bool   `json:"verified_email"`
	ProfilePicture string `json:"picture"`
}

type CountryCode string

const (
	ZA CountryCode = "+27"
)

type PhoneNumber struct {
	Code   CountryCode
	Digits string
}

func (p PhoneNumber) String() string {
	return fmt.Sprintf("%s%s", p.Code, p.Digits)
}
func (p PhoneNumber) IsSouthAfricanPhoneNumber() bool {
	return p.Code == ZA
}
