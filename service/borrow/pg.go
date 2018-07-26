package borrow

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/LuanNgoMinh/example-go/domain"
)

// pgService implmenter for Borrow serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// validate book_id
func BookIdValidate(s *pgService, id domain.UUID) error {
	book := domain.Book{}
	if err := s.db.Where("ID = ?", id).Find(&book).Error; err == nil {
		return nil
	}
	fmt.Printf("BookId invalid")
	return errors.New("Id have not existed yet")
}

// validate user_id
func UserIdValidate(s *pgService, id domain.UUID) error {
	user := domain.User{}
	if err := s.db.Where("ID = ?", id).Find(&user).Error; err == nil {
		return nil
	}
	return errors.New("Id have not existed yet")
}

// book is avaialbe for borrow
func bookIsAvaiable(s *pgService, book_id domain.UUID) error {
	borrow := domain.Borrow{}
	if err := s.db.Where(domain.Borrow{Book_ID: book_id}).Find(&borrow).Error; err == nil {
		return errors.New("Book unavaiable")
	}
	return nil
}

// Create implement Create for Borrow service
func (s *pgService) Create(_ context.Context, p *domain.Borrow) error {
	if err := BookIdValidate(s, p.Book_ID); err != nil {
		return err
	}

	if err := UserIdValidate(s, p.User_ID); err != nil {
		return err
	}

	if err := bookIsAvaiable(s, p.Book_ID); err != nil {
		return err
	}

	return s.db.Create(p).Error
}

// Update implement Update for Borrow service
func (s *pgService) Update(_ context.Context, p *domain.Borrow) (*domain.Borrow, error) {
	old := domain.Borrow{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	old.To = p.To

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Borrow service
func (s *pgService) Find(_ context.Context, p *domain.Borrow) (*domain.Borrow, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Borrow service
func (s *pgService) FindAll(_ context.Context) ([]domain.Borrow, error) {
	res := []domain.Borrow{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Borrow service
func (s *pgService) Delete(_ context.Context, p *domain.Borrow) error {
	old := domain.Borrow{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	return s.db.Delete(old).Error
}
