package repository

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"gorm.io/gorm"
)

type LogProdukRepository struct {
	db *gorm.DB
}

func NewLogProdukRepository(db *gorm.DB) *LogProdukRepository {
	return &LogProdukRepository{db}
}

func (r *LogProdukRepository) Create(log *domain.LogProduk) error {
	return r.db.Create(log).Error
}

func (r *LogProdukRepository) FindByID(id int) (*domain.LogProduk, error) {
	var log domain.LogProduk
	err := r.db.First(&log, id).Error
	return &log, err
}

func (r *LogProdukRepository) FindAll() ([]domain.LogProduk, error) {
	var logs []domain.LogProduk
	err := r.db.Find(&logs).Error
	return logs, err
}