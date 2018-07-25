package book

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/LuanNgoMinh/example-go/domain"
	"github.com/LuanNgoMinh/example-go/service"
)

// CreateData data for CreateBook
type CreateData struct {
	Name        string      `json:"name"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
	Category_id domain.UUID `json:"category_id"`
}

// CreateRequest request struct for CreateBook
type CreateRequest struct {
	Book CreateData `json:"book"`
}

// CreateResponse response struct for CreateBook
type CreateResponse struct {
	Book domain.Book `json:"book"`
}

// StatusCode customstatus code for success create User
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a User
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			book = &domain.Book{
				Name:        req.Book.Name,
				Author:      req.Book.Author,
				Description: req.Book.Description,
				Category_id: req.Book.Category_id,
			}
		)

		err := s.BookService.Create(ctx, book)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Book: *book}, nil
	}
}

// FindRequest request struct for Find a User
type FindRequest struct {
	BookID domain.UUID
}

// FindResponse response struct for Find a User
type FindResponse struct {
	Book *domain.Book `json:"book"`
}

// MakeFindEndPoint make endpoint for find User
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var bookFind domain.Book
		req := request.(FindRequest)
		bookFind.ID = req.BookID

		book, err := s.BookService.Find(ctx, &bookFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Book: book}, nil
	}
}

// FindAllRequest request struct for FindAll User
type FindAllRequest struct{}

// FindAllResponse request struct for find all User
type FindAllResponse struct {
	Books []domain.Book `json:"books"`
}

// MakeFindAllEndpoint make endpoint for find all User
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		books, err := s.BookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Books: books}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID          domain.UUID `json:"-"`
	Name        string      `json:"name"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
	Category_id domain.UUID `json:"category_id"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Book UpdateData `json:"book"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Book domain.Book `json:"book"`
}

// MakeUpdateEndpoint make endpoint for update a User
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			book = domain.Book{
				Model:       domain.Model{ID: req.Book.ID},
				Name:        req.Book.Name,
				Author:      req.Book.Author,
				Description: req.Book.Description,
				Category_id: req.Book.Category_id,
			}
		)

		res, err := s.BookService.Update(ctx, &book)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Book: *res}, nil
	}
}

// DeleteRequest request struct for delete a User
type DeleteRequest struct {
	BookID domain.UUID
}

// DeleteResponse response struct for Find a User
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a User
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			bookFind = domain.Book{}
			req      = request.(DeleteRequest)
		)
		bookFind.ID = req.BookID

		err := s.BookService.Delete(ctx, &bookFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
