package user_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"context"
	"fmt"
)

func (this *UserUseCases) Authorize(ctx context.Context, data *requests.AuthUserRequest) (*responses.UserResponse, error) {
	if !this.deps.UserRdr.IsUserNameExists(ctx, data.UserName) {
		return nil, fmt.Errorf("Username does not exists")
	}

	usr, err := this.deps.UserRdr.GetByUserName(ctx, data.UserName)
	if err != nil {
		return nil, err
	}

	if !usr.VerifyPassword(data.Password) {
		return nil, fmt.Errorf("Uncorrect password")
	}

	return mappers.UserToResponse(usr), nil
}
