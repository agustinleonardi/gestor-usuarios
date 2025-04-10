package middleware

import (
	"net/http"
	"strings"

	"github.com/agustinleonardi/gestor-usuarios/internal/ports"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware(authService ports.AuthService, userRepo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token faltante o inválido"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := authService.VerifyToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		user, err := userRepo.GetByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
			return
		}

		// Setear el usuario en el contexto para acceder desde handlers o middlewares
		c.Set("user", user)

		c.Next()
	}
}
