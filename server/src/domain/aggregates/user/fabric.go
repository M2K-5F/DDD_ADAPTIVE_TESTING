package user

import (
	"adaptivetesting/src/domain/identificators"
	"time"
)

type UserFabric struct{}

func (UserFabric) CreateStudent(username string, password string) (*User, error) {
	usernameVO, err := NewUsername(username)

	if err != nil {
		return nil, err
	}

	hashVO, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		id:           identificators.NewUserID(),
		registeredAt: time.Now(),
		roles:        []Role{StudentRole},
		userName:     usernameVO,
		passwordHash: hashVO,
	}, nil
}

func (UserFabric) CreateTeacher(username string, password string) (*User, error) {
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
		id:           identificators.NewUserID(),
		passwordHash: hash,
	}, nil
}

func (UserFabric) Recover(id string, username string, registeredAt time.Time, passwordHash string, roles []string) (*User, error) {
	usernameVO, err := NewUsername(username)
	if err != nil {
		return nil, err
	}

	passwordVO := Password(passwordHash)

	var rolesVO []Role
	for _, role := range roles {
		roleVO, err := NewRole(role)
		if err != nil {
			return nil, err
		}
		rolesVO = append(rolesVO, roleVO)
	}

	idVO, err := identificators.ParseUserID(id)
	if err != nil {
		return nil, err
	}

	return &User{
		id:           idVO,
		userName:     usernameVO,
		registeredAt: registeredAt,
		passwordHash: passwordVO,
		roles:        rolesVO,
	}, nil
}

var Fabric = UserFabric{}
