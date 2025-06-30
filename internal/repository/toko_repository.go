package repository

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"gorm.io/gorm"
)

type TokoRepository struct {
    db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) *TokoRepository {
    return &TokoRepository{db}
}

func (r *TokoRepository) Create(toko *domain.Toko) error {
    return r.db.Create(toko).Error
}

func (r *TokoRepository) FindByUserID(userID int) (*domain.Toko, error) {
	var toko domain.Toko
	err := r.db.Where("id_user = ?", userID).First(&toko).Error
	return &toko, err
}

func (r *TokoRepository) Update(toko *domain.Toko) error {
	return r.db.Save(toko).Error
}