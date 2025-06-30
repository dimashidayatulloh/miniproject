package repository

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"gorm.io/gorm"
)

type ProdukRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{db}
}

func (r *ProdukRepository) Create(produk *domain.Produk) error {
	return r.db.Create(produk).Error
}

func (r *ProdukRepository) Update(produk *domain.Produk) error {
	return r.db.Save(produk).Error
}

func (r *ProdukRepository) Delete(id int) error {
	return r.db.Delete(&domain.Produk{}, id).Error
}

func (r *ProdukRepository) FindByID(id int) (*domain.Produk, error) {
	var produk domain.Produk
	err := r.db.First(&produk, id).Error
	return &produk, err
}

func (r *ProdukRepository) FindAll() ([]domain.Produk, error) {
	var produks []domain.Produk
	err := r.db.Find(&produks).Error
	return produks, err
}

func (r *ProdukRepository) FindByToko(idToko int) ([]domain.Produk, error) {
	var produks []domain.Produk
	err := r.db.Where("id_toko = ?", idToko).Find(&produks).Error
	return produks, err
}