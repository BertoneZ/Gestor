package dto

// Registro de usuario
type UserRequestRegisterDTO struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"required"` // "admin", "operator"
}

// Login
type UserRequestLoginDTO struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// Respuesta de usuario
type UserResponseDTO struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Role     string `json:"role"`
}