package dto
import "time"
type CreateStockMovementDTO struct {
    ProductID   string `json:"product_id" binding:"required"`
    Type        string `json:"type" binding:"required"` // "in" o "out"
    Quantity    int    `json:"quantity" binding:"required"`
    PerformedBy string `json:"performed_by" binding:"required"`
    Note        string `json:"note,omitempty"`
    Timestamp   string `json:"timestamp" binding:"required"` // formato ISO8601
}
type StockMovementResponseDTO struct {
    ID          string    `json:"id"`
    ProductID   string    `json:"product_id"`
    Type        string    `json:"type"`
    Quantity    int       `json:"quantity"`
    PerformedBy string    `json:"performed_by"`
    Note        string    `json:"note,omitempty"`
    Timestamp   time.Time `json:"timestamp"`
}