package book

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/LuanNgoMinh/example-go/domain"

	"errors"
)

// pgService implmenter for Book serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func bookIsExisted(s *pgService, name string) bool {
	book := domain.Book{}
	if err := s.db.Where(domain.Book{Name: name}).Find(&book).Error; err != nil {
		return false
	}
	return true
}

// Create implement Create for Book service
func (s *pgService) Create(_ context.Context, p *domain.Book) error {
	if !bookIsExisted(s, p.Name) {
		return s.db.Create(p).Error
	} else {
		return errors.New(string(p.Name) + " have already exists")
	}
}

// Update implement Update for Book service
func (s *pgService) Update(_ context.Context, p *domain.Book) (*domain.Book, error) {
	if !bookIsExisted(s, p.Name) {
		return nil, errors.New(string(p.Name) + " have already exists")
	}

	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	old.Name = p.Name

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Book service
func (s *pgService) Find(_ context.Context, p *domain.Book) (*domain.Book, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Book service
func (s *pgService) FindAll(_ context.Context) ([]domain.Book, error) {
	res := []domain.Book{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Book service
func (s *pgService) Delete(_ context.Context, p *domain.Book) error {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	return s.db.Delete(old).Error
}
