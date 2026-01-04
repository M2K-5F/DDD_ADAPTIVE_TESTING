package user_usercases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type UserRegistration struct {
	userRepo user.IUserRepository
}

func (r *UserRegistration) Execute(data requests.RegisterUserDTO) (*responses.UserResponse, error) {
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

func (uc *UserRegistration) createUserByRole(data requests.RegisterUserDTO) (*user.User, error) {
	switch data.Role {
	case requests.Teacher:
		return user.NewTeacher(data.UserName, data.Password)
	case requests.Student:
		return user.NewStudent(data.UserName, data.Password)
	default:
		return nil, fmt.Errorf("unknown role: %s", data.Role)
	}
}

func (UserRegistration) Fabric(userRepoImplementation user.IUserRepository) *UserRegistration {
	return &UserRegistration{
		userRepo: userRepoImplementation,
	}
}
