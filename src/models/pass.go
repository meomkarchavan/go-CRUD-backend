package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pass struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PassId       string             `json:"passid"`
	UserId       string             `json:"userid"`
	UserFullName string             `json:"userFullName"`
	PurposeId    string             `json:"purposeid"`
	Date         primitive.DateTime `json:"date"`
	Approved     bool               `json:"approved"`
	Rejected     bool               `json:"rejected"`
}
