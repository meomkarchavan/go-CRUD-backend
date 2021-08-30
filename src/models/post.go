package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	PostId   string             `json:"postid"`
	Title    string             `json:"title" validate:"gte=5,lte=30"`
	Content  string             `json:"content" validate:"gte=5,lte=30"`
	UserId   string             `json:"userid"`
	Likes    int64              `json:"likes"`
	ImageUrl string             `json:"imageurl"`
}
