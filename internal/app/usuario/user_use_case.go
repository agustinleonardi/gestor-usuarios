package app

import (
	"errors"
	"time"

	"github.com/agustinleonardi/gestor-usuarios/internal/domain/user"
	"github.com/agustinleonardi/gestor-usuarios/internal/ports"
)

type UserUseCase struct {
	repo        ports.UserRepository
	authService ports.AuthService
}

func NewUserUseCase(repo ports.UserRepository, auth ports.AuthService) *UserUseCase {
	return &UserUseCase{
		repo:        repo,
		authService: auth,
	}
}

func (s *UserUseCase) Register(name, email, password string) error {
	existingUser, err := s.repo.GetByEmail(email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("el correo ya esta registrado")
	}
	hashedPassword, err := s.authService.HashPassword(password)
	if err != nil {
		return err
	}

	return s.repo.Create(&user.User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
}

func (s *UserUseCase) List() ([]*user.User, error) {
	return s.repo.List()
}
func (s *UserUseCase) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("credenciales invalidas")
	}
	if err := s.authService.CheckPassword(password, user.Password); err != nil {
		return "", errors.New("credenciales invalidas")
	}
	token, err := s.authService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
