package users

import (
	"errors"
	"fmt"
)

func (u *users.User) MakeFriend(target *User) (*User, error) {
	if u == target {
		return nil, errors.New(fmt.Sprintf("can't friend with himself: %v", u))
	}
	u.Friends = append(u.Friends, target)
	return u, nil
}
