package commands

import (
	"adaptivetesting/src/domain/aggregates/user"
	"context"

	"github.com/jackc/pgx/v5"
)

func SaveUser(ctx context.Context, txn pgx.Tx, user *user.User) error {
	return nil
}
