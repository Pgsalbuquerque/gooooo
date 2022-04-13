package dto

type UserRequestDTO struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	CPF      string `bson:"cpf,omitempty"`
}
