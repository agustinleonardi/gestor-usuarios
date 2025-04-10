package db

import (
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/user"
	"github.com/agustinleonardi/gestor-usuarios/internal/ports"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// Constructor para inyectar GORM
func NewGormUserRepository(db *gorm.DB) ports.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) GetByEmail(email string) (*user.User, error) {
	var user user.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // no existe el usuario
		}
		return nil, err // error de consulta
	}
	return &user, nil
}
func (r *GormUserRepository) GetByID(id int) (*user.User, error) {
	var user user.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Update(user *user.User) error {
	return nil
}

func (r *GormUserRepository) Delete(id int) error {
	return nil
}

func (r *GormUserRepository) List() ([]*user.User, error) {
	var users []*user.User
	if err := r.db.Preload("Roles").Preload("Tokens").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
