package ports

import (
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/role"
)

type RoleRepository interface {
	GetByID(id int) (*role.Role, error)
	GetByName(name string) (*role.Role, error)
	List() ([]role.Role, error)
	AssignRoleToUser(userID, roleID int) error
}
