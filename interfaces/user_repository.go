package interfaces

import (
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/models"
)

type IUserRepository interface {
	Store(p models.User) (dto.UserResponseDTO, error)
	Update(p models.User) (dto.UserResponseDTO, error)
	GetById(id string) (models.User, error)
	Delete(id string) error
}
