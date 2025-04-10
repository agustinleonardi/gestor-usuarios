package ports

import "github.com/agustinleonardi/gestor-usuarios/internal/domain/token"

type TokenRepository interface {
	Create(token *token.Token) error
	GetByToken(value string) (*token.Token, error)
	DeleteExpired() error
	DeleteByUserID(userID int) error
}
