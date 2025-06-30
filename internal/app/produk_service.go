package app

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

type ProdukService struct {
	repo *repository.ProdukRepository
}

func NewProdukService(repo *repository.ProdukRepository) *ProdukService {
	return &ProdukService{repo}
}

func (s *ProdukService) CreateProduk(input *domain.Produk) error {
	return s.repo.Create(input)
}

func (s *ProdukService) UpdateProduk(id int, input *domain.Produk) error {
	produk, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	// Update fields
	produk.NamaProduk = input.NamaProduk
	produk.Slug = input.Slug
	produk.HargaReseller = input.HargaReseller
	produk.HargaKonsumen = input.HargaKonsumen
	produk.Stok = input.Stok
	produk.Deskripsi = input.Deskripsi
	produk.IdToko = input.IdToko
	produk.IdCategory = input.IdCategory
	return s.repo.Update(produk)
}

func (s *ProdukService) DeleteProduk(id int) error {
	return s.repo.Delete(id)
}

func (s *ProdukService) GetProdukByID(id int) (*domain.Produk, error) {
	return s.repo.FindByID(id)
}

func (s *ProdukService) GetAllProduk() ([]domain.Produk, error) {
	return s.repo.FindAll()
}

func (s *ProdukService) GetProdukByToko(idToko int) ([]domain.Produk, error) {
	return s.repo.FindByToko(idToko)
}