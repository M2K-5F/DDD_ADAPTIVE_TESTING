package queries

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/identificators"
	"adaptivetesting/src/infrastructure/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseReader struct {
	pool *pgxpool.Pool
}

func (this *CourseReader) GetByID(
	ctx context.Context,
	courseID identificators.CourseID,
) (
	*course.Course,
	error,
) {
	var rows models.CourseRows
	err := pgxscan.Get(ctx, this.pool, &rows, `
		select c.id, c.created_by_id, c.name, c.is_archived
		from courses c 
		where c.id = $1
		limit 1;
	`, courseID)

	if err != nil {
		return nil, err
	}
	return course.Fabric.Recover(
		rows.ID,
		rows.CreatedByID,
		rows.Name,
		rows.IsArchived,
	)
}

func (this *CourseReader) GetByName(ctx context.Context, name string) (*course.Course, error) {
	var rows models.CourseRows
	err := pgxscan.Get(ctx, this.pool, &rows, `
		select c.id, c.created_by_id, c.name, c.is_archived
		from courses c 
		where c.name = $1
		limit 1;
	`, name)

	if err != nil {
		return nil, err
	}
	return course.Fabric.Recover(
		rows.ID,
		rows.CreatedByID,
		rows.Name,
		rows.IsArchived,
	)
}

func (this *CourseReader) SelectCreatedByUser(ctx context.Context, userID identificators.UserID) ([]*course.Course, error) {
	var rows []models.CourseRows
	err := pgxscan.Select(ctx, this.pool, &rows, `
		select c.id, c.created_by_id, c.name, c.is_archived
		from courses c 
		where c.created_by_id = $1
	`)

	if err != nil {
		return nil, err
	}

	var courses []*course.Course
	for _, row := range rows {
		crs, err := course.Fabric.Recover(
			row.ID,
			row.CreatedByID,
			row.Name,
			row.IsArchived,
		)
		if err != nil {
			return nil, err
		}

		courses = append(courses, crs)
	}
	return courses, nil
}
