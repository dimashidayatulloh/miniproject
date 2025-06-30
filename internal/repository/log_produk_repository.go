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

func (r *LogProdukRepository) FindAllPaginatedFiltered(page, limit int, jenis, keterangan string, idProduk int) ([]domain.LogProduk, int64, error) {
	var logs []domain.LogProduk
	var total int64

	db := r.db.Model(&domain.LogProduk{})

	if jenis != "" {
		db = db.Where("jenis LIKE ?", "%"+jenis+"%")
	}
	if keterangan != "" {
		db = db.Where("keterangan LIKE ?", "%"+keterangan+"%")
	}
	if idProduk > 0 {
		db = db.Where("id_produk = ?", idProduk)
	}

	// Hitung total sesuai filter
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Limit(limit).Offset(offset).Order("id desc").Find(&logs).Error
	return logs, total, err
}