package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string             `json:"userid"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email" validate:"email"`
	Username  string             `json:"username" validate:"alphanum,required,gte=5,lte=20"`
	Password  string             `json:"password" validate:"gte=5,lte=100"`
}
