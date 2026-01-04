package user

import "adaptivetesting/src/domain/aggregates/identificators"

type IUserRepository interface {
	Create(user *User) error
	GetByID(id identificators.UserID) (*User, error)
	Update(user *User) error
	IsUsernameExists(username string) bool
	GetByUserName(username string) (*User, error)
}
