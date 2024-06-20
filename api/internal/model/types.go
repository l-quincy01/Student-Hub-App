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

func (pn PhoneNumber) String() string {
	return fmt.Sprintf("%s%s", pn.Code, pn.Digits)
}
