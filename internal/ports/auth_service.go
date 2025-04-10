package ports

type AuthService interface {
	GenerateToken(user int) (string, error)
	VerifyToken(token string) (int, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
