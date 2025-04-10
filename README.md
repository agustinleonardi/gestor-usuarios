# 🧑‍💻 Sistema de Gestión de Usuarios – Go + Hexagonal Architecture

Este proyecto es una API REST desarrollada en **Go**, que implementa un sistema completo de autenticación, registro y gestión de usuarios, roles y permisos. Está construido siguiendo la **arquitectura hexagonal**, con separación clara de capas, ideales para escalar, testear y mantener el sistema a largo plazo.

---

## ⚙️ Tecnologías utilizadas

- **Go** (Golang) – lenguaje principal
- **Gin** – framework HTTP
- **GORM** – ORM para bases de datos relacionales
- **MySQL** – base de datos (en Azure)
- **bcrypt** – para hash de contraseñas
- **JWT** – autenticación segura con tokens
- **Arquitectura hexagonal** – separación de dominio, puertos y adaptadores

---

## 📦 Funcionalidades

### ✅ Gestión de usuarios

- Registro de nuevos usuarios
- Validación de email único
- Hash de contraseñas con bcrypt

### ✅ Seguridad

- Autenticación con JWT (en proceso)
- Middleware para proteger rutas (pendiente)
- Verificación de tokens

### ✅ Roles y permisos

- Creación de roles (ej: `user`, `admin`)
- Asignación automática del rol `"user"` al registrarse
- Relaciones muchos a muchos: `user_roles`, `role_permissions`

---

## 🧱 Arquitectura del proyecto

/internal
├── /domain → modelos del negocio (User, Role, Token, etc.)
├── /ports → interfaces (UserRepository, AuthService, etc.)
├── /app → lógica del negocio (casos de uso)
├── /adapters │ ├── /db → conexión con la base (GORM)
│ ├── /auth → seguridad: bcrypt, JWT
│ └── /http → controladores y rutas (Gin)
