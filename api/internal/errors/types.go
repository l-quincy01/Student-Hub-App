package errors

import (
	"fmt"
	"net/http"
)

type ApiError interface {
	Error() string
	StatusCode() int
}

type InternalServerError struct {
	Err error
}

func (i InternalServerError) Error() string {
	return fmt.Sprintf("An unexpected error occured. status: %d, err: %v", i.StatusCode(), i.Err)
}

func (i InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}

type UnathorizedError struct {
	Err error
}

func (u UnathorizedError) Error() string {
	return fmt.Sprintf("Not authorized to perform operation. status: %d, err: %v", u.StatusCode(), u.Err)
}

func (u UnathorizedError) StatusCode() int {
	return http.StatusUnauthorized
}

type BadRequestError struct {
	Err error
}

func (b BadRequestError) Error() string {
	return fmt.Sprintf("Unable to process the operation. status: %d, err: %v", b.StatusCode(), b.Err)
}

func (b BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}
