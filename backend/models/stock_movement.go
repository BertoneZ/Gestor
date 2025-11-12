package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type StockMovement struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ProductID  primitive.ObjectID `bson:"product_id" json:"product_id"`
    Type       string             `bson:"type" json:"type"` // "in" o "out"
    Quantity   int                `bson:"quantity" json:"quantity"`
    PerformedBy string            `bson:"performed_by" json:"performed_by"` // nombre o ID de usuario
    Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
    Note       string             `bson:"note,omitempty" json:"note,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}