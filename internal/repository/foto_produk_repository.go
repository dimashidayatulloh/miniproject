package repository

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"gorm.io/gorm"
)

type FotoProdukRepository struct {
	db *gorm.DB
}

func NewFotoProdukRepository(db *gorm.DB) *FotoProdukRepository {
	return &FotoProdukRepository{db}
}

func (r *FotoProdukRepository) Create(foto *domain.FotoProduk) error {
	return r.db.Create(foto).Error
}

func (r *FotoProdukRepository) FindByID(id int) (*domain.FotoProduk, error) {
	var foto domain.FotoProduk
	err := r.db.First(&foto, id).Error
	return &foto, err
}

func (r *FotoProdukRepository) FindAllByProduk(idProduk int) ([]domain.FotoProduk, error) {
	var fotos []domain.FotoProduk
	err := r.db.Where("id_produk = ?", idProduk).Find(&fotos).Error
	return fotos, err
}