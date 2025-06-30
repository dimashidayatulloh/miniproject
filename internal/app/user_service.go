package app

import (
	"errors"

	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Register(user *domain.User) error {
	// Hash password sebelum simpan ke DB
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.KataSandi = string(hashed)
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (*domain.User, error) {
    user, err := s.repo.FindByEmail(email)
    if err != nil {
        return nil, errors.New("email tidak ditemukan")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(password)); err != nil {
        return nil, errors.New("password salah")
    }
    return user, nil
}

func (s *UserService) GetByID(id int) (*domain.User, error) {
	return s.repo.GetByID(id)
}