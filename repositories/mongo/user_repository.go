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

	_, err := db.Database("user").Collection("users").UpdateOne(ctx, bson.M{"_id": p.ID}, bson.D{{"$set", bson.D{{"password", p.Password}}}})

	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	dto := dto.UserResponseDTO{
		Username: p.Username,
		CPF:      p.CPF,
	}

	return dto, nil
}

func (s *UserRepository) GetById(ID string) (models.User, error) {
	db := database.GetDB()
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	objectId, err := primitive.ObjectIDFromHex(ID)
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

func (s *UserRepository) Delete(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	db := database.GetDB()

	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	_, err = db.Database("user").Collection("users").DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil

}

func (s *UserRepository) GetByCPF(CPF string) (models.User, error) {
	db := database.GetDB()
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	result := models.User{}

	err := db.Database("user").Collection("users").FindOne(ctx, bson.M{"cpf": CPF}).Decode(&result)
	if err != nil {
		return models.User{}, err
	}

	return result, nil

}

func (s *UserRepository) ChangePlan(ID string, Plan string) error {
	db := database.GetDB()
	ctx, close := context.WithTimeout(context.TODO(), 10*time.Second)
	defer close()

	objectId, _ := primitive.ObjectIDFromHex(ID)

	_, err := db.Database("user").Collection("users").UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{"$set", bson.D{{"plan", Plan}}}})
	if err != nil {
		return err
	}

	return nil
}
