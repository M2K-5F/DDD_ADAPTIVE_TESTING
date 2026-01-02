package question

import (
	"adaptivetesting/src/domain/aggregates/topic"

	"github.com/google/uuid"
)

type QuestionID uuid.UUID

type Question struct {
	id      QuestionID
	topicID topic.TopicID
	text    QuestionText
	answers []Answer
}
