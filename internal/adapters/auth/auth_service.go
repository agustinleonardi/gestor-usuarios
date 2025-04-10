package auth

import (
	"errors"
	"time"

	"github.com/agustinleonardi/gestor-usuarios/internal/ports"
	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("clave-secreta-supersegura")

type BcryptAuthService struct {
}

func NewBcryptAuthService() ports.AuthService {
	return &BcryptAuthService{}
}

// Hashear contraseña
func (s *BcryptAuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verificar contraseña ingresada contra hash guardado
func (s *BcryptAuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *BcryptAuthService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *BcryptAuthService) VerifyToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("token inválido")
	}

	// Extraer claims y user_id
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("user_id inválido en el token")
		}
		return int(userIDFloat), nil
	}

	return 0, errors.New("no se pudieron leer los claims")
}
