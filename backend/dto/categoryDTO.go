
package dto

type CreateCategoryDTO struct {
    Name string `json:"name" binding:"required"`
}
type CategoryResponseDTO struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}