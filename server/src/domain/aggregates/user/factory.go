package user

import (
	"time"

	"github.com/google/uuid"
)

func NewStudent(username string, password string) (*User, error) {
	usernameVO, err := NewUsername(username)

	if err != nil {
		return nil, err
	}

	hashVO, err := NewPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		id:           UserID(uuid.New()),
		registeredAt: time.Now(),
		roles:        []Role{StudentRole},
		userName:     usernameVO,
		passwordHash: hashVO,
	}, nil
}

func NewTeacher(username string, password string) (*User, error) {
	usernameVO, err := NewUsername(username)
	if err != nil {
		return nil, err
	}

	hash, err := NewPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		userName:     usernameVO,
		registeredAt: time.Now(),
		roles:        []Role{TeacherRole, StudentRole},
		id:           UserID(uuid.New()),
		passwordHash: hash,
	}, nil
}
