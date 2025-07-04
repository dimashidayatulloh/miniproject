package repository

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"gorm.io/gorm"
)

type AlamatRepository struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) *AlamatRepository {
	return &AlamatRepository{db}
}

func (r *AlamatRepository) Create(alamat *domain.Alamat) error {
	return r.db.Create(alamat).Error
}

func (r *AlamatRepository) Update(alamat *domain.Alamat) error {
	return r.db.Save(alamat).Error
}

func (r *AlamatRepository) Delete(id int, userID int) error {
	return r.db.Where("id = ? AND id_user = ?", id, userID).Delete(&domain.Alamat{}).Error
}

func (r *AlamatRepository) FindByID(id int, userID int) (*domain.Alamat, error) {
	var alamat domain.Alamat
	err := r.db.Where("id = ? AND id_user = ?", id, userID).First(&alamat).Error
	return &alamat, err
}

func (r *AlamatRepository) FindAllByUser(userID int) ([]domain.Alamat, error) {
	var alamat []domain.Alamat
	err := r.db.Where("id_user = ?", userID).Find(&alamat).Error
	return alamat, err
}

func (r *AlamatRepository) FindAllByUserPaginatedFiltered(userID, page, limit int, nama, judul string) ([]domain.Alamat, int64, error) {
	var alamat []domain.Alamat
	var total int64

	db := r.db.Model(&domain.Alamat{}).Where("id_user = ?", userID)
	if nama != "" {
		db = db.Where("nama_penerima LIKE ?", "%"+nama+"%")
	}
	if judul != "" {
		db = db.Where("judul_alamat LIKE ?", "%"+judul+"%")
	}

	// Hitung total sesuai filter
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Limit(limit).Offset(offset).Find(&alamat).Error
	return alamat, total, err
}