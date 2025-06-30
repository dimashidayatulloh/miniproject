package app

import (
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

type AlamatService struct {
	repo *repository.AlamatRepository
}

func NewAlamatService(repo *repository.AlamatRepository) *AlamatService {
	return &AlamatService{repo}
}

func (s *AlamatService) CreateAlamat(userID int, input *domain.Alamat) error {
	input.IdUser = userID
	return s.repo.Create(input)
}

func (s *AlamatService) UpdateAlamat(userID int, id int, input *domain.Alamat) error {
	alamat, err := s.repo.FindByID(id, userID)
	if err != nil {
		return err
	}
	alamat.JudulAlamat = input.JudulAlamat
	alamat.NamaPenerima = input.NamaPenerima
	alamat.NoTelp = input.NoTelp
	alamat.DetailAlamat = input.DetailAlamat
	return s.repo.Update(alamat)
}

func (s *AlamatService) DeleteAlamat(userID int, id int) error {
	return s.repo.Delete(id, userID)
}

func (s *AlamatService) GetAlamatByID(userID int, id int) (*domain.Alamat, error) {
	return s.repo.FindByID(id, userID)
}

func (s *AlamatService) GetAllAlamatByUser(userID int) ([]domain.Alamat, error) {
	return s.repo.FindAllByUser(userID)
}