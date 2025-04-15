package handlers

import (
	"net/http"

	app "github.com/agustinleonardi/gestor-usuarios/internal/app/usuario"
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUseCase *app.UserUseCase
}

func NewAuthHandler(userUseCase *app.UserUseCase) *AuthHandler {
	return &AuthHandler{userUseCase: userUseCase}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos invlidos"})
		return
	}

	token, err := h.userUseCase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales invalidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
func (h *AuthHandler) Me(c *gin.Context) {
	userRaw, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	user, ok := userRaw.(*user.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno de tipo de usuario"})
		return
	}

	c.JSON(http.StatusOK, user)
}
