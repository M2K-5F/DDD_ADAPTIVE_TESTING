package user_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"context"
	"fmt"
)

type UserAuthorization struct {
	userReader interfaces.IUserReader
}

func (this *UserAuthorization) Execute(ctx context.Context, data *requests.AuthUserDTO) (*responses.UserResponse, error) {
	if !this.userReader.IsUserNameExists(ctx, data.UserName) {
		return nil, fmt.Errorf("Username does not exists")
	}

	usr, err := this.userReader.GetByUserName(ctx, data.UserName)
	if err != nil {
		return nil, err
	}

	if !usr.VerifyPassword(data.Password) {
		return nil, fmt.Errorf("Uncorrect password")
	}

	return mappers.UserToResponse(usr), nil
}

func FabricAuthorization(userReader interfaces.IUserReader) *UserAuthorization {
	return &UserAuthorization{
		userReader: userReader,
	}
}
