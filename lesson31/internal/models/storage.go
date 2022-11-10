package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Storage interface {
	Get(id primitive.ObjectID) *User
	Save(u *User) (*User, error)
	AllUsers() []User
	Delete(id primitive.ObjectID) error
	Update(u *User) (*User, error)
	MakeFriend(source, target *User) (*User, error)
	DeleteFriend(source, target *User) (*User, error)
	GetFriends(id primitive.ObjectID) ([]User, error)
}
