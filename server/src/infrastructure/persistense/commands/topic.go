package commands

import (
	"adaptivetesting/src/domain/aggregates/topic"
	"context"

	"github.com/jackc/pgx/v5"
)

func SaveTopic(ctx context.Context, txn pgx.Tx, course *topic.Topic) error {
	return nil
}
