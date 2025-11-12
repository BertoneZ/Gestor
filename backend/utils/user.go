package utils

import (
	"time"
	"github.com/BertoneZ/Gestor.git/backend/dto"
	"github.com/BertoneZ/Gestor.git/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RequestToUser(dto dto.UserRequestRegisterDTO, hashedPassword string) models.User {
	return models.User{
		ID:        primitive.NewObjectID(),
		Username:  dto.Username,
		Email:     dto.Email,
		Password:  hashedPassword,
		Role:      dto.Role,
		CreatedAt: time.Now(),
	}
}

func UserToResponse(u models.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}
}