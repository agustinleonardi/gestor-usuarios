package main

import (
	"log"

	"github.com/agustinleonardi/gestor-usuarios/internal/adapters/auth"
	"github.com/agustinleonardi/gestor-usuarios/internal/adapters/db"
	"github.com/agustinleonardi/gestor-usuarios/internal/adapters/http/handlers"
	"github.com/agustinleonardi/gestor-usuarios/internal/adapters/http/middleware"
	app "github.com/agustinleonardi/gestor-usuarios/internal/app/usuario"
	"github.com/agustinleonardi/gestor-usuarios/internal/domain"
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/permission"
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/role"
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/token"
	"github.com/agustinleonardi/gestor-usuarios/internal/domain/user"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 📦 Conexión a la base
	dsn := "root:Quilmesagustin8@tcp(127.0.0.1:3306)/gestor_usuarios?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	err = database.AutoMigrate(&user.User{}, &role.Role{}, &permission.Permission{}, &token.Token{}, &domain.RolePermission{}, &domain.UserRole{})
	if err != nil {
		log.Fatalf("Error al migrar modelos: %v", err)
	}

	log.Println("Conectado a la base de datos y modelos migrados.")

	userRepo := db.NewGormUserRepository(database)
	authService := auth.NewBcryptAuthService()

	//casos de uso
	userUseCase := app.NewUserUseCase(userRepo, authService)

	//handlers
	userHandler := handlers.NewUserHandler(userUseCase)
	authHandler := handlers.NewAuthHandler(userUseCase)

	//middleware

	// Router
	r := gin.New()
	userGroup := r.Group("users")
	authGroup := r.Group("auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.GET("/me", middleware.JWTMiddleware(authService, userRepo), authHandler.Me)
	}

	userGroup.POST("/register", userHandler.Register)
	userGroup.GET("/", userHandler.ListUsers)

	// Iniciar servidor
	r.Run(":8081")
}
