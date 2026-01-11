package interfaces

import (
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/group"
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
	"context"
)

type IWriter interface {
	Execute(ctx context.Context, fn func(writer ITransactionWriter) error) error
}

type ITransactionWriter interface {
	SaveUser(*user.User) error
	SaveCourse(*course.Course) error
	SaveQuestion(*question.Question) error
	SaveGroup(*group.Group) error
	SaveTopic(*topic.Topic) error
}
