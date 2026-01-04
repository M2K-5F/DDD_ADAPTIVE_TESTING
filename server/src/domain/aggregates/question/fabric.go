package question

import (
	"adaptivetesting/src/domain/aggregates/identificators"
	"adaptivetesting/src/domain/aggregates/topic"

	"github.com/google/uuid"
)

type AnswerDTO struct {
	Text      string
	IsCorrect bool
}

func NewQuestion(topic topic.Topic, text string, answers []AnswerDTO) (*Question, error) {
	var answerList []Answer

	for i, answerDTO := range answers {
		answer, err := NewAnswer(i+1, answerDTO.Text, answerDTO.IsCorrect)

		if err != nil {
			return nil, err
		}

		answerList = append(
			answerList,
			*answer,
		)
	}

	textVO, err := NewQuestionText(text)

	if err != nil {
		return nil, err
	}

	return &Question{
		id:      identificators.QuestionID(uuid.New()),
		topicID: topic.ID(),
		text:    textVO,
		answers: answerList,
	}, nil
}
