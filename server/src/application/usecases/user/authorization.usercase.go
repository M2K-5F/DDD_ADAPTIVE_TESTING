package user_usercases

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/application/mappers"
	user "adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type AuthorizationUserUC struct {
	userRepo user.IUserRepository
}

func (uc *AuthorizationUserUC) Execute(data *dto.AuthUserDTO) (*dto.UserResponse, error) {
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
