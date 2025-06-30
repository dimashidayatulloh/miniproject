package app

import (
	"errors"

	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) CreateCategory(isAdmin bool, input *domain.Category) error {
	if !isAdmin {
		return errors.New("hanya admin yang bisa membuat kategori")
	}
	return s.repo.Create(input)
}

func (s *CategoryService) UpdateCategory(isAdmin bool, id int, input *domain.Category) error {
	if !isAdmin {
		return errors.New("hanya admin yang bisa update kategori")
	}
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	cat.NamaCategory = input.NamaCategory
	return s.repo.Update(cat)
}

func (s *CategoryService) DeleteCategory(isAdmin bool, id int) error {
	if !isAdmin {
		return errors.New("hanya admin yang bisa hapus kategori")
	}
	return s.repo.Delete(id)
}

func (s *CategoryService) GetAllCategory() ([]domain.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) GetCategoryByID(id int) (*domain.Category, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) GetAllCategoryPaginatedFiltered(page, limit int, nama string) ([]domain.Category, int64, error) {
	return s.repo.FindAllPaginatedFiltered(page, limit, nama)
}