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