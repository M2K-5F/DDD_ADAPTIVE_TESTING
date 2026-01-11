package interfaces

import (
	"adaptivetesting/src/domain/aggregates/user"
	"adaptivetesting/src/domain/identificators"
	"context"
)

type IUserReader interface {
	GetByID(ctx context.Context, id identificators.UserID) (*user.User, error)
	GetByUserName(ctx context.Context, username string) (*user.User, error)
	IsUserNameExists(ctx context.Context, username string) bool
}
