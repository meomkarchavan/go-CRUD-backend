package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TodoId      string             `json:"todoid"`
	Title       string             `json:"title" validate:"gte=5,lte=30"`
	Description string             `json:"description" validate:"gte=5,lte=30"`
	UserId      string             `json:"userid"`
	Done        bool               `json:"done"`
}
