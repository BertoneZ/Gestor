package dto 
type CreateProductDTO struct {
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
    Category    string  `json:"category" binding:"required"`
    Price       float64 `json:"price" binding:"required"`
    Stock       int     `json:"stock" binding:"required"`
}
type UpdateProductDTO struct {
    Description *string  `json:"description,omitempty"`
    Category    *string  `json:"category,omitempty"`
    Price       *float64 `json:"price,omitempty"`
    Stock       *int     `json:"stock,omitempty"`   
}
type ProductResponseDTO struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Category    string  `json:"category"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    Active      bool    `json:"active"`
}