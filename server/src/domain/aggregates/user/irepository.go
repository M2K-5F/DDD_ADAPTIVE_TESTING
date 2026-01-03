package user

type IUserRepository interface {
	Create(user *User) error
	GetByID(id UserID) (*User, error)
	Update(user *User) error
	IsUsernameExists(username string) bool
	GetByUserName(username string) (*User, error)
}
