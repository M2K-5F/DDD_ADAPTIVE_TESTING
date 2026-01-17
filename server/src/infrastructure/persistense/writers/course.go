package writers

import (
	"adaptivetesting/src/domain/aggregates/course"
	"context"

	"github.com/jackc/pgx/v5"
)

func SaveCourse(ctx context.Context, txn pgx.Tx, course *course.Course) error {
	model := course.ToPersistense()

	_, err := txn.Exec(ctx, `
		insert into 
		courses (id, created_by_id, name, is_archived)
		values ($1, $2, $3, $4)
		on conflict (id) 
		do update set 
			name = excluded.name,
			created_by_id = excluded.created_by_id,
			is_archived = excluded.is_archived
		`,
		model.ID,
		model.CreatedByID,
		model.Name,
		model.IsArchived,
	)
	if err != nil {
		return err
	}

	return nil
}
