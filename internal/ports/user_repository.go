package ports

import (
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/user"
)

type UserRepository interface {
	Create(user *user.User) error
	GetByID(id int) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Update(user *user.User) error
	Delete(id int) error
	List() ([]*user.User, error)
}
