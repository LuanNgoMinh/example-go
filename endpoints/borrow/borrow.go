package borrow

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/LuanNgoMinh/example-go/domain"
	"github.com/LuanNgoMinh/example-go/service"
)

// CreateData data for CreateBorrow
type CreateData struct {
	Book_ID domain.UUID `json:"book_id"`
	User_ID domain.UUID `json:"user_id"`
	To      time.Time   `json:"to`
}

// CreateRequest request struct for CreateBorrow
type CreateRequest struct {
	Borrow CreateData `json:"borrow"`
}

// CreateResponse response struct for CreateBorrow
type CreateResponse struct {
	Borrow domain.Borrow `json:"borrow"`
}

// StatusCode customstatus code for success create User
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a User
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req    = request.(CreateRequest)
			borrow = &domain.Borrow{
				Book_ID: req.Borrow.Book_ID,
				User_ID: req.Borrow.User_ID,
				To:      req.Borrow.To,
			}
		)

		err := s.BorrowService.Create(ctx, borrow)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Borrow: *borrow}, nil
	}
}

// FindRequest request struct for Find a User
type FindRequest struct {
	BorrowID domain.UUID
}

// FindResponse response struct for Find a User
type FindResponse struct {
	Borrow *domain.Borrow `json:"borrow"`
}

// MakeFindEndPoint make endpoint for find User
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var borrowFind domain.Borrow
		req := request.(FindRequest)
		borrowFind.ID = req.BorrowID

		borrow, err := s.BorrowService.Find(ctx, &borrowFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Borrow: borrow}, nil
	}
}

// FindAllRequest request struct for FindAll User
type FindAllRequest struct{}

// FindAllResponse request struct for find all User
type FindAllResponse struct {
	Borrows []domain.Borrow `json:"borrow"`
}

// MakeFindAllEndpoint make endpoint for find all User
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		borrow, err := s.BorrowService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Borrows: borrow}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID      domain.UUID `json:"-"`
	Book_ID domain.UUID `json:"book_id"`
	User_ID domain.UUID `json:"user_id"`
	From    time.Time   `json:"from"`
	To      time.Time   `json:"to`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Borrow UpdateData `json:"borrow"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Borrow domain.Borrow `json:"borrow"`
}

// MakeUpdateEndpoint make endpoint for update a User
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			book = domain.Borrow{
				Model:   domain.Model{ID: req.Borrow.ID},
				Book_ID: req.Borrow.Book_ID,
				User_ID: req.Borrow.User_ID,
				To:      req.Borrow.To,
			}
		)

		res, err := s.BorrowService.Update(ctx, &book)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Borrow: *res}, nil
	}
}

// DeleteRequest request struct for delete a User
type DeleteRequest struct {
	BorrowID domain.UUID
}

// DeleteResponse response struct for Find a User
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a User
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			borrowFind = domain.Borrow{}
			req        = request.(DeleteRequest)
		)
		borrowFind.ID = req.BorrowID

		err := s.BorrowService.Delete(ctx, &borrowFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
