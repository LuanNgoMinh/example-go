package borrow

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/LuanNgoMinh/example-go/domain"
	borrowEndpoint "github.com/LuanNgoMinh/example-go/endpoints/borrow"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	borrowID, err := domain.UUIDFromString(chi.URLParam(r, "borrow_id"))
	if err != nil {
		return nil, err
	}
	return borrowEndpoint.FindRequest{BorrowID: borrowID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return borrowEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req borrowEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	borrowID, err := domain.UUIDFromString(chi.URLParam(r, "borrow_id"))
	if err != nil {
		return nil, err
	}

	var req borrowEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Borrow.ID = borrowID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	borrowID, err := domain.UUIDFromString(chi.URLParam(r, "borrow_id"))
	if err != nil {
		return nil, err
	}
	return borrowEndpoint.DeleteRequest{BorrowID: borrowID}, nil
}
