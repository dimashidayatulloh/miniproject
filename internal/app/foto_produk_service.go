package app

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

type FotoProdukService struct {
	repo *repository.FotoProdukRepository
}

func NewFotoProdukService(repo *repository.FotoProdukRepository) *FotoProdukService {
	return &FotoProdukService{repo}
}

func (s *FotoProdukService) CreateFotoProduk(input *domain.FotoProduk) error {
	return s.repo.Create(input)
}

func (s *FotoProdukService) GetFotoProdukByID(id int) (*domain.FotoProduk, error) {
	return s.repo.FindByID(id)
}

func (s *FotoProdukService) GetAllFotoProdukByProduk(idProduk int) ([]domain.FotoProduk, error) {
	return s.repo.FindAllByProduk(idProduk)
}

func (s *FotoProdukService) GetAllFotoProdukByProdukPaginatedFiltered(idProduk, page, limit int, url string) ([]domain.FotoProduk, int64, error) {
	return s.repo.FindAllByProdukPaginatedFiltered(idProduk, page, limit, url)
}