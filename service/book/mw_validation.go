package book

import (
	"context"

	"github.com/LuanNgoMinh/example-go/domain"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware ...
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, book *domain.Book) (err error) {
	if book.Name == "" {
		return ErrNameIsRequire
	}

	if len(book.Name) < 5 {
		return ErrNameNotEnoughLength
	}

	if book.Description == "" {
		return ErrDescriptionIsRequire
	}

	if len(book.Description) < 5 {
		return ErrDescriptionNotEnoughLength
	}

	return mw.Service.Create(ctx, book)
}

func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Book, error) {
	return mw.Service.FindAll(ctx)
}

func (mw validationMiddleware) Find(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	return mw.Service.Find(ctx, book)
}

func (mw validationMiddleware) Update(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	if book.Name == "" {
		return nil, ErrNameIsRequire
	}

	if len(book.Name) < 5 {
		return nil, ErrNameNotEnoughLength
	}

	if book.Description == "" {
		return nil, ErrDescriptionIsRequire
	}

	if len(book.Description) < 5 {
		return nil, ErrDescriptionNotEnoughLength
	}

	return mw.Service.Update(ctx, book)
}

func (mw validationMiddleware) Delete(ctx context.Context, book *domain.Book) error {
	return mw.Service.Delete(ctx, book)
}
