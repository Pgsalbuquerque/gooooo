package sessions

import (
	"errors"
	"fmt"
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/interfaces"
	"strateegy/user-service/services/encrypt"
)

type SessionService struct {
	userRepository interfaces.IUserRepository
}

func NewSessionService(r interfaces.IUserRepository) *SessionService {
	return &SessionService{
		userRepository: r,
	}
}

func (s *SessionService) Login(data dto.UserLoginDTO) (TokenResponse, error) {
	user, err := s.userRepository.GetByCPF(data.CPF)
	if err != nil {
		return TokenResponse{}, err
	}

	if encrypt.SHA256Encoder(data.Password) != user.Password {
		fmt.Println(data.Password, user.Password)
		return TokenResponse{}, errors.New("Password does not match")
	}

	token, err := NewJWTService().GenerateToken(user.ID.Hex())
	if err != nil {
		return TokenResponse{}, errors.New("Internal Error")
	}

	return TokenResponse{
		Token: token,
	}, nil
}
