package dto

type UserDeleteDTO struct {
	Password        string `bson:"password,omitempty"`
	ConfirmPassword string `bson:"confirmPassword,omitempty"`
}
