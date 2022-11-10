package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name    string             `json:"name"`
	Age     int                `json:"age"`
	Friends []*User            `json:"friends"`
}

type UserInRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type FriendRequest struct {
	Source_id primitive.ObjectID `bson:"_id" json:"source_id,omitempty"`
	Target_id primitive.ObjectID `bson:"_id" json:"target_id,omitempty"`
}
