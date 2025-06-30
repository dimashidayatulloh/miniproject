package repository

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id int) error {
	return r.db.Delete(&domain.Category{}, id).Error
}

func (r *CategoryRepository) FindByID(id int) (*domain.Category, error) {
	var cat domain.Category
	err := r.db.First(&cat, id).Error
	return &cat, err
}

func (r *CategoryRepository) FindAll() ([]domain.Category, error) {
	var cats []domain.Category
	err := r.db.Find(&cats).Error
	return cats, err
}