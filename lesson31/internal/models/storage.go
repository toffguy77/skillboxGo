package models

type Storage interface {
	Get(id string) *User
	Save(u *User) (*User, error)
	AllUsers() []User
	Delete(id string) error
	Update(u *User) (*User, error)
	MakeFriend(source, target *User) (*User, error)
	DeleteFriend(source, target *User) (*User, error)
	GetFriends(id string) ([]User, error)
}
