package user_usercases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type UserAuthorization struct {
	userRepo user.IUserRepository
}

func (uc *UserAuthorization) Execute(data *requests.AuthUserDTO) (*responses.UserResponse, error) {
	if !uc.userRepo.IsUsernameExists(data.UserName) {
		return nil, fmt.Errorf("Username does not exists")
	}

	usr, err := uc.userRepo.GetByUserName(data.UserName)
	if err != nil {
		return nil, err
	}

	if !usr.VerifyPassword(data.Password) {
		return nil, fmt.Errorf("Uncorrect password")
	}

	return mappers.UserToDTO(usr), nil
}
