package writers

import (
	"adaptivetesting/src/domain/aggregates/topic"
	"context"

	"github.com/jackc/pgx/v5"
)

func SaveTopic(ctx context.Context, txn pgx.Tx, topic *topic.Topic) error {
	model := topic.ToPersistense()

	_, err := txn.Exec(ctx, `
		insert into 
		topics (id, by_course_id, created_by_id, name, is_archived)
		values ($1, $2, $3, $4, $5)
		on conflict (id)
		do update set 
			by_course_id = excluded.by_course_id,
			created_by_id = excluded.created_by_id,
			name = excluded.name,
			is_archived = excluded.is_archived
		`,
		model.ID,
		model.ByCourseID,
		model.CreatedByID,
		model.Name,
		model.IsArchived,
	)
	if err != nil {
		return err
	}
	return nil
}
