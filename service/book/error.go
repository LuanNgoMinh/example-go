package book

import "net/http"

var (
	ErrNameIsRequire              = errNameIsRequire{}
	ErrNameNotEnoughLength        = errNameNotEnoughLength{}
	ErrDescriptionIsRequire       = errDescriptionIsRequire{}
	ErrDescriptionNotEnoughLength = errDescriptionNotEnoughLength{}
)

type errNameIsRequire struct{}

func (errNameIsRequire) Error() string {
	return "Name is require"
}

func (errNameIsRequire) StatusCode() int {
	return http.StatusBadRequest
}

type errNameNotEnoughLength struct{}

func (errNameNotEnoughLength) Error() string {
	return "Name is not enough length"
}

func (errNameNotEnoughLength) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionIsRequire struct{}

func (errDescriptionIsRequire) Error() string {
	return "Description is require"
}

func (errDescriptionIsRequire) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionNotEnoughLength struct{}

func (errDescriptionNotEnoughLength) Error() string {
	return "Description must have length greater than 5"
}

func (errDescriptionNotEnoughLength) StatusCode() int {
	return http.StatusBadRequest
}
