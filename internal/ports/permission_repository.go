package ports

import (
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/permission"
)

type PermissionRepository interface {
	GetByID(id int) (*permission.Permission, error)
	ListByRoleID(roleID int) ([]permission.Permission, error)
}
