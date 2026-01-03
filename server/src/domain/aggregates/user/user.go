package user

import (
	"time"

	"github.com/google/uuid"
)

// USER
type UserID uuid.UUID

func (i UserID) String() string {
	return uuid.UUID(i).String()
}

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

func (u *User) UserName() Username {
	return u.userName
}

func (u *User) Roles() []Role {
	return u.roles
}

func (u *User) RegisteredAt() time.Time {
	return u.registeredAt
}

func (u *User) VerifyPassword(plain string) bool {
	return u.passwordHash.Verify(plain)
}
