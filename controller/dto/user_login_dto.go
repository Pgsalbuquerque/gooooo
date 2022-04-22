package dto

type UserLoginDTO struct {
	CPF      string `bson:"CPF,omitempty"`
	Password string `bson:"password,omitempty"`
}
