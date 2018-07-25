package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound       = errNotFound{}
	ErrUnknown        = errUnknown{}
	ErrNameIsRequired = errNameIsRequired{}
	ErrNameLength     = errNameLength{}
	ErrNameUnique     = errNameUnique{}
	ErrNameExisted    = errNameExisted{}
	ErrRecordNotFound = errRecordNotFound{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameLength struct{}

func (errNameLength) Error() string {
	return "user name is not enough length"
}

func (errNameLength) StatusCode() int {
	return http.StatusBadRequest
}

type errNameUnique struct{}

func (errNameUnique) Error() string {
	return "Name have already existed"
}

func (errNameUnique) StatusCode() int {
	return http.StatusBadRequest
}

type errNameExisted struct{}

func (errNameExisted) Error() string {
	return "Name have not already existed"
}

func (errNameExisted) StatusCode() int {
	return http.StatusNotFound
}
