package dto

type UserUpdatePasswordDTO struct {
	Password        string `bson:"password,omitempty"`
	ConfirmPassword string `bson:"confirmPassword,omitempty"`
	LastPassword    string `bson:"lastPassword,omitempty"`
}
