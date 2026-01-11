package mappers

import (
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/domain/aggregates/user"
)

func UserToResponse(user *user.User) *responses.UserResponse {
	roles := []string{}
	for _, role := range user.Roles() {
		roles = append(roles, role.String())
	}

	return &responses.UserResponse{
		UserName:     user.UserName().String(),
		RegisteredAt: user.RegisteredAt(),
		Roles:        roles,
		ID:           user.ID().String(),
	}
}
