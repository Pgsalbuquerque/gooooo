package interfaces

import (
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/models"
)

type IUserRepository interface {
	Store(p models.User) (dto.UserResponseDTO, error)
	Update(p models.User) (dto.UserResponseDTO, error)
	GetById(ID string) (models.User, error)
	Delete(ID string) error
	GetByCPF(CPF string) (models.User, error)
	ChangePlan(ID string, Plan string) error
}
