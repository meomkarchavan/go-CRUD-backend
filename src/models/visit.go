package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Visit struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty"`
	VisitId      string              `json:"visitid"`
	UserId       string              `json:"userid"`
	UserFullName string              `json:"userFullName"`
	PurposeId    string              `json:"purposeid"`
	Date         primitive.Timestamp `json:"date"`
	Approved     bool                `json:"approved"`
	Rejected     bool                `json:"rejected"`
}
