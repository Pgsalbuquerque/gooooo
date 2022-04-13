package user

import (
	"errors"
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/interfaces"
	"strateegy/user-service/models"
	"strateegy/user-service/models/plan"
)

type UserService struct {
	repository interfaces.IUserRepository
}

func (s *UserService) CreateUser(data dto.UserRequestDTO) (dto.UserResponseDTO, error) {
	user := models.User{
		Username: data.Username,
		Password: data.Password,
		CPF:      data.CPF,
		Plan:     plan.BASIC,
	}

	return s.repository.Store(user)

}

func NewUserService(r interfaces.IUserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) UpdateUserPassword(data dto.UserUpdatePasswordDTO) (dto.UserResponseDTO, error) {
	user, err := s.repository.GetById(data.ID)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	if user.Password != data.LastPassword {
		return dto.UserResponseDTO{}, errors.New("Last password does not match")
	}

	if data.Password != data.ConfirmPassword {
		return dto.UserResponseDTO{}, errors.New("Passwords does not match")
	}

	updatedUser := models.User{
		ID:       user.ID,
		Username: user.Username,
		Password: data.Password,
		CPF:      user.CPF,
		Plan:     user.Plan,
	}

	result, err := s.repository.Update(updatedUser)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	return result, nil
}
