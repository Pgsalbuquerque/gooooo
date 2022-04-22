package user

import (
	"errors"
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/interfaces"
	"strateegy/user-service/models"
	"strateegy/user-service/models/plan"
	"strateegy/user-service/services/encrypt"
	"strateegy/user-service/services/sessions"
)

type UserService struct {
	repository interfaces.IUserRepository
}

func NewUserService(r interfaces.IUserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) CreateUser(data dto.UserRequestDTO) (dto.UserResponseDTO, error) {
	password := encrypt.SHA256Encoder(data.Password)

	user := models.User{
		Username: data.Username,
		Password: password,
		CPF:      data.CPF,
		Plan:     plan.BASIC,
	}

	return s.repository.Store(user)

}

func (s *UserService) UpdateUserPassword(Token string, data dto.UserUpdatePasswordDTO) (dto.UserResponseDTO, error) {
	ID, err := sessions.NewJWTService().GetIDFromToken(Token)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	user, err := s.repository.GetById(ID)
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

func (s *UserService) DeleteUser(token string, data dto.UserDeleteDTO) error {
	ID, err := sessions.NewJWTService().GetIDFromToken(token)
	if err != nil {
		return err
	}

	user, err := s.repository.GetById(ID)
	if err != nil {
		return err
	}

	if data.Password != data.ConfirmPassword {
		return errors.New("Password does not match")
	}

	if user.Password != data.Password {
		return errors.New("Password incorrect")
	}

	err = s.repository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(token string) (models.User, error) {
	ID, err := sessions.NewJWTService().GetIDFromToken(token)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.repository.GetById(ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// func (s *UserService) UpdatePlan(token string, plan string) error {

// }
