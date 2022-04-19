package dto

type UserDeleteDTO struct {
	ID              string `bson:"id,ID,omitempty"`
	Password        string `bson:"password,omitempty"`
	ConfirmPassword string `bson:"confirmPassword,omitempty"`
}
