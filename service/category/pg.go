package category

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/LuanNgoMinh/example-go/domain"

	"errors"
)

// pgService implmenter for Category serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Check category name is unique
func CategoryIsUnique(s *pgService, name string) error {
	category := domain.Category{}
	if err := s.db.Where(domain.Category{Name: name}).Find(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
	}

	return nil
}

// Create implement Create for Category service
func (s *pgService) Create(_ context.Context, p *domain.Category) error {
	if err := CategoryIsUnique(s, p.Name); err != nil {
		return s.db.Create(p).Error
	}

	return errors.New("Category name has already existed")
}

// Update implement Update for Category service
func (s *pgService) Update(_ context.Context, p *domain.Category) (*domain.Category, error) {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	old.Name = p.Name

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Category service
func (s *pgService) Find(_ context.Context, p *domain.Category) (*domain.Category, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Category service
func (s *pgService) FindAll(_ context.Context) ([]domain.Category, error) {
	res := []domain.Category{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Category service
// a2ab5d7a-9aa5-4ec4-afee-6962d559ad68
func (s *pgService) Delete(_ context.Context, p *domain.Category) error {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	// fmt.Println(p.ID)
	defer DeleteAllBooksBelong(s, p.ID)

	return s.db.Delete(old).Error
}

// Delete all book belong category with id
func DeleteAllBooksBelong(s *pgService, id domain.UUID) error {
	if err := s.db.Where("category_id = ?", id).Delete(&domain.Book{}); err != nil {
		return nil
	}
	return errors.New("Delete book error")
}
