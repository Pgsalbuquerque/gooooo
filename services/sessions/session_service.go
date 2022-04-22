package sessions

import (
	"context"
	"errors"
	"strateegy/user-service/controller/dto"
	"strateegy/user-service/grpc"
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
		return TokenResponse{}, errors.New("Password does not match")
	}

	conn := grpc.GetConn()
	client := grpc.NewSendTokenClient(conn)

	req := &grpc.ID{
		ID: user.ID.Hex(),
	}

	res, err := client.RequestToken(context.Background(), req)
	if err != nil {
		return TokenResponse{}, errors.New("Internal Error")
	}

	token := res.GetToken()
	if err != nil {
		return TokenResponse{}, errors.New("Internal Error")
	}

	return TokenResponse{
		Token: token,
	}, nil
}
