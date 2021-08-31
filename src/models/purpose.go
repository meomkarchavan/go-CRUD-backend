package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Purpose struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	PurposeId string             `json:"purposeid"`
	Title     string             `json:"title"`
}
