
package dto

// Crear categoría
type CategoryRequestCreateDTO struct {
    Name string `json:"name" binding:"required"`
}

// Respuesta de categoría
type CategoryResponseDTO struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}