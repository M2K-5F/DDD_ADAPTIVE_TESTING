package commands

import (
	"adaptivetesting/src/domain/aggregates/question"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SaveQuestion(ctx context.Context, txn pgx.Tx, question *question.Question) error {
	return fmt.Errorf("")
}
