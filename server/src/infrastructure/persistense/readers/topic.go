package readers

import (
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/identificators"
	"adaptivetesting/src/infrastructure/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TopicReader struct {
	pool *pgxpool.Pool
}

func (this *TopicReader) GetByID(ctx context.Context, topicID identificators.TopicID) (*topic.Topic, error) {
	var rows models.TopicRows
	if err := pgxscan.Get(ctx, this.pool, &rows, `
		select t.id, t.name, t.created_by_id, t.by_course_id, t.is_archived from topics
		where t.id = $1
		limit 1;
	`); err != nil {
		return nil, err
	}
	return topic.Fabric.Recover(
		rows.Name,
		rows.ByCourseID,
		rows.CreatedByID,
		rows.ID,
		rows.IsArchived,
	)
}
