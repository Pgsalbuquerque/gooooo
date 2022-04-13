package dto

type UserResponseDTO struct {
	Username string `bson:"username,omitempty"`
	CPF      string `bson:"cpf,omitempty"`
}
