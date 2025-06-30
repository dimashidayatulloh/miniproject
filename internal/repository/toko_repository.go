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

func (r *TokoRepository) FindAllPaginatedFiltered(page, limit int, nama, urlFoto string, userID int) ([]domain.Toko, int64, error) {
	var tokos []domain.Toko
	var total int64

	db := r.db.Model(&domain.Toko{})
	if nama != "" {
		db = db.Where("nama_toko LIKE ?", "%"+nama+"%")
	}
	if urlFoto != "" {
		db = db.Where("url_foto LIKE ?", "%"+urlFoto+"%")
	}
	if userID > 0 {
		db = db.Where("id_user = ?", userID)
	}

	// Hitung total sesuai filter
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Limit(limit).Offset(offset).Find(&tokos).Error
	return tokos, total, err
}