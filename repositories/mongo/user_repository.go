package mongorepo

import (
	"context"
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/database"
	"strateegy/user-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
}

func (s *UserRepository) Store(p models.User) (dto.UserResponseDTO, error) {
	db := database.GetDB()
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	_, err := db.Database("user").Collection("users").InsertOne(ctx, &p)

	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	dto := dto.UserResponseDTO{
		Username: p.Username,
		CPF:      p.CPF,
	}

	return dto, nil
}

func (s *UserRepository) Update(p models.User) (dto.UserResponseDTO, error) {
	db := database.GetDB()
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	_, err := db.Database("user").Collection("users").UpdateByID(ctx, p.ID, &p)

	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	dto := dto.UserResponseDTO{
		Username: p.Username,
		CPF:      p.CPF,
	}

	return dto, nil
}

func (s *UserRepository) GetById(id string) (models.User, error) {
	db := database.GetDB()
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	result := models.User{}

	err = db.Database("user").Collection("users").FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
	if err != nil {
		return models.User{}, err
	}

	return result, nil

}
