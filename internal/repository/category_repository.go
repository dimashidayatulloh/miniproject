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

func (r *CategoryRepository) FindAllPaginatedFiltered(page, limit int, nama string) ([]domain.Category, int64, error) {
	var cats []domain.Category
	var total int64

	db := r.db.Model(&domain.Category{})

	if nama != "" {
		db = db.Where("nama_category LIKE ?", "%"+nama+"%")
	}

	// Hitung total sesuai filter
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Limit(limit).Offset(offset).Find(&cats).Error
	return cats, total, err
}