package commands

import (
	"adaptivetesting/src/domain/aggregates/course"
	"context"

	"github.com/jackc/pgx/v5"
)

func SaveCourse(ctx context.Context, txn pgx.Tx, course *course.Course) error {
	return nil
}
