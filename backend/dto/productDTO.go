package dto 
// Crear producto
type ProductRequestCreateDTO struct {
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
    Category    string  `json:"category" binding:"required"`
    Price       float64 `json:"price" binding:"required"`
    Stock       int     `json:"stock" binding:"required"`
}

// Actualizar producto
type ProductRequestUpdateDTO struct {
    Description *string  `json:"description,omitempty"`
    Category    *string  `json:"category,omitempty"`
    Price       *float64 `json:"price,omitempty"`
    Stock       *int     `json:"stock,omitempty"`
    Active      *bool    `json:"active,omitempty"`
}

// Respuesta de producto
type ProductResponseDTO struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Category    string  `json:"category"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    Active      bool    `json:"active"`
}