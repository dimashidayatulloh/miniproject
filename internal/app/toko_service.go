package app

import (
	"errors"

	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

type TokoService struct {
	repo *repository.TokoRepository
}

func NewTokoService(repo *repository.TokoRepository) *TokoService {
	return &TokoService{repo}
}

// Membuat toko (dipanggil otomatis saat register user)
func (s *TokoService) CreateToko(toko *domain.Toko) error {
	return s.repo.Create(toko)
}

// Mendapatkan toko milik user sendiri
func (s *TokoService) GetTokoByUserID(userID int) (*domain.Toko, error) {
	return s.repo.FindByUserID(userID)
}

// Update toko milik user sendiri
func (s *TokoService) UpdateToko(userID int, update *domain.Toko) error {
	toko, err := s.repo.FindByUserID(userID)
	if err != nil {
		return err
	}
	// Validasi toko harus milik user
	if toko.IdUser != userID {
		return errors.New("unauthorized")
	}
	toko.NamaToko = update.NamaToko
	toko.UrlFoto = update.UrlFoto
	return s.repo.Update(toko)
}

func (s *TokoService) GetAllTokoPaginatedFiltered(page, limit int, nama, urlFoto string, userID int) ([]domain.Toko, int64, error) {
	return s.repo.FindAllPaginatedFiltered(page, limit, nama, urlFoto, userID)
}