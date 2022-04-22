package user

import (
	"context"
	"errors"
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/grpc"
	"strateegy/user-service/interfaces"
	"strateegy/user-service/models"
	"strateegy/user-service/models/plan"
	"strateegy/user-service/services/encrypt"
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

func (s *UserService) UpdateUserPassword(token string, data dto.UserUpdatePasswordDTO) (dto.UserResponseDTO, error) {

	conn := grpc.GetConn()
	client := grpc.NewSendIDClient(conn)

	req := &grpc.Token{
		Token: token,
	}

	res, err := client.RequestID(context.Background(), req)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	ID := res.GetID()

	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	user, err := s.repository.GetById(ID)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	if user.Password != encrypt.SHA256Encoder(data.LastPassword) {
		return dto.UserResponseDTO{}, errors.New("Last password does not match")
	}

	if data.Password != data.ConfirmPassword {
		return dto.UserResponseDTO{}, errors.New("Passwords does not match")
	}

	updatedUser := models.User{
		ID:       user.ID,
		Username: user.Username,
		Password: encrypt.SHA256Encoder(data.Password),
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
	conn := grpc.GetConn()
	client := grpc.NewSendIDClient(conn)

	req := &grpc.Token{
		Token: token,
	}

	res, err := client.RequestID(context.Background(), req)
	if err != nil {
		return err
	}

	ID := res.GetID()

	user, err := s.repository.GetById(ID)
	if err != nil {
		return err
	}

	if data.Password != data.ConfirmPassword {
		return errors.New("Password does not match")
	}

	if user.Password != encrypt.SHA256Encoder(data.Password) {
		return errors.New("Password incorrect")
	}

	err = s.repository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(token string) (models.User, error) {

	conn := grpc.GetConn()
	client := grpc.NewSendIDClient(conn)

	req := &grpc.Token{
		Token: token,
	}

	res, err := client.RequestID(context.Background(), req)
	if err != nil {
		return models.User{}, err
	}

	ID := res.GetID()

	user, err := s.repository.GetById(ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// func (s *UserService) UpdatePlan(token string, plan string) error {

// }
