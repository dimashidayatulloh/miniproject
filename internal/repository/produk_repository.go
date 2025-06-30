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

func (r *ProdukRepository) FindAllPaginated(page int, limit int) ([]domain.Produk, int64, error) {
	var produks []domain.Produk
	var total int64

	// Hitung total data
	if err := r.db.Model(&domain.Produk{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := r.db.Limit(limit).Offset(offset).Find(&produks).Error
	return produks, total, err
}

func (r *ProdukRepository) FindAllFiltered(page, limit int, nama, kategori string, hargaMin, hargaMax int) ([]domain.Produk, int64, error) {
    var produks []domain.Produk
    var total int64

    db := r.db.Model(&domain.Produk{})

    if nama != "" {
        db = db.Where("nama_produk LIKE ?", "%"+nama+"%")
    }
    if kategori != "" {
        db = db.Where("id_category = ?", kategori) // atau join ke kategori jika ingin nama
    }
    if hargaMin > 0 {
        db = db.Where("harga_konsumen >= ?", hargaMin)
    }
    if hargaMax > 0 {
        db = db.Where("harga_konsumen <= ?", hargaMax)
    }

    // Hitung total data sesuai filter
    if err := db.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    offset := (page - 1) * limit
    err := db.Limit(limit).Offset(offset).Find(&produks).Error
    return produks, total, err
}