package app

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

type LogProdukService struct {
	repo *repository.LogProdukRepository
}

func NewLogProdukService(repo *repository.LogProdukRepository) *LogProdukService {
	return &LogProdukService{repo}
}

func (s *LogProdukService) CreateLogProduk(input *domain.LogProduk) error {
	return s.repo.Create(input)
}

func (s *LogProdukService) GetLogProdukByID(id int) (*domain.LogProduk, error) {
	return s.repo.FindByID(id)
}

func (s *LogProdukService) GetAllLogProduk() ([]domain.LogProduk, error) {
	return s.repo.FindAll()
}

func (s *LogProdukService) GetAllLogProdukPaginatedFiltered(page, limit int, jenis, keterangan string, idProduk int) ([]domain.LogProduk, int64, error) {
	return s.repo.FindAllPaginatedFiltered(page, limit, jenis, keterangan, idProduk)
}