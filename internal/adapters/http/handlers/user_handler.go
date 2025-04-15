package handlers

import (
	"net/http"

	app "github.com/agustinleonardi/gestor-usuarios/internal/app/usuario"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase *app.UserUseCase
}

func NewUserHandler(userUseCase *app.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := h.userUseCase.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado con éxito"})
}

// GET /users
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.userUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}
