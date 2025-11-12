package models
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Username string             `bson:"username" json:"username"`
    Email    string             `bson:"email" json:"email"`
    Password string             `bson:"password" json:"-"` // nunca se expone
    Role     string             `bson:"role" json:"role"` // "admin", "operator"
    CreatedAt time.Time         `bson:"created_at" json:"created_at"`
    DeletedAt *time.Time        `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}