package user

import (
	"time"

	"github.com/google/uuid"
)

// USER
type UserID uuid.UUID

type User struct {
	id           UserID
	registeredAt time.Time
	userName     Username
	roles        []Role
	passwordHash Password
}

func (this *User) ID() UserID {
	return this.id
}

func (u *User) IsStudent() bool {
	for _, role := range u.roles {
		if role == StudentRole {
			return true
		}
	}
	return false
}

func (u *User) IsTeacher() bool {
	for _, role := range u.roles {
		if role == TeacherRole {
			return true
		}
	}
	return false
}
