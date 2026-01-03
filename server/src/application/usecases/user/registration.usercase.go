package user_usercases

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type RegisterUserUC struct {
	userRepo user.IUserRepository
}

func (r *RegisterUserUC) Execute(data dto.RegisterUserDTO) (*dto.UserResponse, error) {
	if r.userRepo.IsUsernameExists(data.UserName) {
		return nil, fmt.Errorf("Username already exists")
	}

	newUser, err := r.createUserByRole(data)

	if err != nil {
		return nil, err
	}

	if err := r.userRepo.Create(newUser); err != nil {
		return nil, err
	}

	return mappers.UserToDTO(newUser), nil
}

func (uc *RegisterUserUC) createUserByRole(data dto.RegisterUserDTO) (*user.User, error) {
	switch data.Role {
	case dto.Teacher:
		return user.NewTeacher(data.UserName, data.Password)
	case dto.Student:
		return user.NewStudent(data.UserName, data.Password)
	default:
		return nil, fmt.Errorf("unknown role: %s", data.Role)
	}
}

func (RegisterUserUC) Fabric(userRepoImplementation user.IUserRepository) *RegisterUserUC {
	return &RegisterUserUC{
		userRepo: userRepoImplementation,
	}
}
