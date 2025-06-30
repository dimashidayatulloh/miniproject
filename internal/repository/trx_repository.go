package repository

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"gorm.io/gorm"
)

type TrxRepository struct {
	db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) *TrxRepository {
	return &TrxRepository{db}
}

func (r *TrxRepository) Create(trx *domain.Trx, details []domain.DetailTrx) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(trx).Error; err != nil {
			return err
		}
		for i := range details {
			details[i].IdTrx = trx.ID // set foreign key
			if err := tx.Create(&details[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *TrxRepository) FindByUser(userID int) ([]domain.Trx, error) {
	var trxs []domain.Trx
	err := r.db.Where("id_user = ?", userID).Find(&trxs).Error
	return trxs, err
}

func (r *TrxRepository) FindByID(id int, userID int) (*domain.Trx, error) {
	var trx domain.Trx
	err := r.db.Where("id = ? AND id_user = ?", id, userID).First(&trx).Error
	return &trx, err
}

func (r *TrxRepository) FindDetailsByTrx(trxID int) ([]domain.DetailTrx, error) {
	var details []domain.DetailTrx
	err := r.db.Where("id_trx = ?", trxID).Find(&details).Error
	return details, err
}

func (r *TrxRepository) FindByUserPaginatedFiltered(userID, page, limit int, kodeInvoice, metode, tanggal string, minTotal, maxTotal int) ([]domain.Trx, int64, error) {
	var trxs []domain.Trx
	var total int64

	db := r.db.Model(&domain.Trx{}).Where("id_user = ?", userID)

	if kodeInvoice != "" {
		db = db.Where("kode_invoice LIKE ?", "%"+kodeInvoice+"%")
	}
	if metode != "" {
		db = db.Where("method_bayar LIKE ?", "%"+metode+"%")
	}
	if tanggal != "" {
		db = db.Where("DATE(created_at) = ?", tanggal)
	}
	if minTotal > 0 {
		db = db.Where("harga_total >= ?", minTotal)
	}
	if maxTotal > 0 {
		db = db.Where("harga_total <= ?", maxTotal)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Order("created_at desc").Limit(limit).Offset(offset).Find(&trxs).Error
	return trxs, total, err
}