package mappers

import (
	"adaptivetesting/src/application/dto"
	"adaptivetesting/src/domain/aggregates/user"
)

func UserToDTO(user *user.User) *dto.UserResponse {
	roles := []string{}
	for _, role := range user.Roles() {
		roles = append(roles, role.String())
	}

	return &dto.UserResponse{
		UserName:     user.UserName().String(),
		RegisteredAt: user.RegisteredAt(),
		Roles:        roles,
		ID:           user.ID().String(),
	}
}
