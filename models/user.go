package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"Username,omitempty"`
	Password string             `bson:"password,omitempty"`
	CPF      string             `bson:"cpf,omitempty"`
	Plan     string             `bson:"plan,omitempty"`
}
