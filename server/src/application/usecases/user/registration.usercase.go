package user_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"adaptivetesting/src/domain/aggregates/user"
	"context"
	"fmt"
)

type UserRegistration struct {
	writer     interfaces.IWriter
	userReader interfaces.IUserReader
}

func (this *UserRegistration) Execute(ctx context.Context, data requests.RegisterUserDTO) (*responses.UserResponse, error) {
	if this.userReader.IsUserNameExists(ctx, data.UserName) {
		return nil, fmt.Errorf("Username already exists")
	}

	newUser, err := this.createUserByRole(data)

	if err != nil {
		return nil, err
	}

	if err := this.writer.Execute(ctx, func(writer interfaces.ITransactionWriter) error {
		return writer.SaveUser(newUser)
	}); err != nil {
		return nil, err
	}

	return mappers.UserToResponse(newUser), nil
}

func (uc *UserRegistration) createUserByRole(data requests.RegisterUserDTO) (*user.User, error) {
	switch data.Role {
	case requests.Teacher:
		return user.Fabric.CreateTeacher(data.UserName, data.Password)
	case requests.Student:
		return user.Fabric.CreateStudent(data.UserName, data.Password)
	default:
		return nil, fmt.Errorf("unknown role: %s", data.Role)
	}
}
