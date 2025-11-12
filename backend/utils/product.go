package utils

import (
	"time"
	"github.com/BertoneZ/Gestor.git/backend/dto"
	"github.com/BertoneZ/Gestor.git/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func RequestToProduct(dto dto.ProductRequestCreateDTO) models.Product {
	return models.Product{
		ID:          primitive.NewObjectID(),
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
		Price:       dto.Price,
		Stock:       dto.Stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Active:      true,
	}
}

func ProductToResponse(p models.Product) dto.ProductResponseDTO {
	return dto.ProductResponseDTO{
		ID:          p.ID.Hex(),
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		Active:      p.Active,
	}
}