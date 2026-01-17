package writers

import (
	"adaptivetesting/src/domain/aggregates/group"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SaveGroup(ctx context.Context, txn pgx.Tx, group *group.Group) error {
	return fmt.Errorf("")
}
