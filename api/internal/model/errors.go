package model

import (
	"fmt"
	"net/http"
)

type InternalServerError struct {
	Err error
}

func (i InternalServerError) Error() string {
	return fmt.Sprintf("An unexpected error occured. status: %d, err: %v", http.StatusInternalServerError, i.Err)
}

type UnathorizedError struct {
	Err error
}

func (u UnathorizedError) Error() string {
	return fmt.Sprintf("Not authorized to perform operation. status: %d, err: %v", http.StatusUnauthorized, u.Err)
}

type BadRequestError struct {
	Err error
}

func (b BadRequestError) Error() string {
	return fmt.Sprintf("Unable to process the operation. status: %d, err: %v", http.StatusBadRequest, b.Err)
}
