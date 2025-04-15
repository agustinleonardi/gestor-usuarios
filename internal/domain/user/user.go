package user

import (
	"time"

	"github.com/agustinleonardi/gestor-usuarios/internal/domain/role"
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/token"
)

type User struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"type:varchar(100);unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Roles  []role.Role `gorm:"many2many:user_roles"`
	Tokens []token.Token
}
