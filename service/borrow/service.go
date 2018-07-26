package borrow

import (
	"context"

	"github.com/LuanNgoMinh/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Borrow) error
	Update(ctx context.Context, p *domain.Borrow) (*domain.Borrow, error)
	Find(ctx context.Context, p *domain.Borrow) (*domain.Borrow, error)
	FindAll(ctx context.Context) ([]domain.Borrow, error)
	Delete(ctx context.Context, p *domain.Borrow) error
}
