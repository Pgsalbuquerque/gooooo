package dto

type UserUpdatePasswordDTO struct {
	ID              string `bson:"id,omitempty"`
	Password        string `bson:"password,omitempty"`
	ConfirmPassword string `bson:"confirmPassword,omitempty"`
	LastPassword    string `bson:"lastPassword,omitempty"`
}
