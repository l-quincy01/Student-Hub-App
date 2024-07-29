package model

import "fmt"

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
