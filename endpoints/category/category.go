package category

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/LuanNgoMinh/example-go/domain"
	"github.com/LuanNgoMinh/example-go/service"
)

// CreateData data for CreateCategory
type CreateData struct {
	Name string `json:"name"`
}

// CreateRequest request struct for CreateCategory
type CreateRequest struct {
	Category CreateData `json:"Category"`
}

// CreateResponse response struct for CreateCategory
type CreateResponse struct {
	Category domain.Category `json:"Category"`
}

// StatusCode customstatus code for success create Category
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Category
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			Category = &domain.Category{
				Name: req.Category.Name,
			}
		)

		err := s.CategoryService.Create(ctx, Category)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Category: *Category}, nil
	}
}

// FindRequest request struct for Find a Category
type FindRequest struct {
	CategoryID domain.UUID
}

// FindResponse response struct for Find a Category
type FindResponse struct {
	Category *domain.Category `json:"Category"`
}

// MakeFindEndPoint make endpoint for find Category
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var CategoryFind domain.Category
		req := request.(FindRequest)
		CategoryFind.ID = req.CategoryID

		Category, err := s.CategoryService.Find(ctx, &CategoryFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Category: Category}, nil
	}
}

// FindAllRequest request struct for FindAll Category
type FindAllRequest struct{}

// FindAllResponse request struct for find all Category
type FindAllResponse struct {
	Categorys []domain.Category `json:"Categorys"`
}

// MakeFindAllEndpoint make endpoint for find all Category
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		Categorys, err := s.CategoryService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Categorys: Categorys}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID   domain.UUID `json:"-"`
	Name string      `json:"name"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Category UpdateData `json:"Category"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Category domain.Category `json:"Category"`
}

// MakeUpdateEndpoint make endpoint for update a Category
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(UpdateRequest)
			Category = domain.Category{
				Model: domain.Model{ID: req.Category.ID},
				Name:  req.Category.Name,
			}
		)

		res, err := s.CategoryService.Update(ctx, &Category)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Category: *res}, nil
	}
}

// DeleteRequest request struct for delete a Category
type DeleteRequest struct {
	CategoryID domain.UUID
}

// DeleteResponse response struct for Find a Category
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a Category
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			CategoryFind = domain.Category{}
			req          = request.(DeleteRequest)
		)
		CategoryFind.ID = req.CategoryID

		err := s.CategoryService.Delete(ctx, &CategoryFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
