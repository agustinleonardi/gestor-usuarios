package role

import "github.com/agustinleonardi/gestor-usuarios/internal/domain/permission"

type Role struct {
	ID          int                     `gorm:"primaryKey"`
	Name        string                  `gorm:"not null;unique"`
	Permissions []permission.Permission `gorm:"many2many:role_permissions"`
}
