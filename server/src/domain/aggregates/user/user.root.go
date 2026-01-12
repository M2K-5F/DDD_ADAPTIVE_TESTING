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

func (this *User) ID() identificators.UserID {
	return this.id
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

type UserPersistense struct {
	ID           string
	RegisteredAt time.Time
	UserName     string
	Roles        []string
	PasswordHash string
}

func (this *User) ToPersistense() *UserPersistense {
	var persistenseRoles []string
	for _, role := range this.roles {
		persistenseRoles = append(persistenseRoles, role.String())
	}

	return &UserPersistense{
		ID:           this.id.String(),
		RegisteredAt: this.registeredAt,
		UserName:     this.userName.String(),
		Roles:        persistenseRoles,
		PasswordHash: string(this.passwordHash),
	}
}
