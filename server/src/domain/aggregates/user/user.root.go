package user

import (
	"adaptivetesting/src/domain/identificators"
	"time"
)

type User struct {
	id           identificators.UserID
	registeredAt time.Time
	userName     Username
	roles        []Role
	passwordHash Password
}

// GETTERS
func (this *User) ID() identificators.UserID {
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
