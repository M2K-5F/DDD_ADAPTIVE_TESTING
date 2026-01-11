package question

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"

	"github.com/google/uuid"
)

type AnswerCreate struct {
	IsCorrect bool
	Text      string
}

type QuestionFabric struct{}

func (this QuestionFabric) CreateQuestion(topic identificators.TopicID, text string, answers []*AnswerCreate) (*Question, error) {

	textVO, err := NewQuestionText(text)

	if err != nil {
		return nil, err
	}

	var answerEntities []*Answer
	for i, answerCreate := range answers {
		answerEntity, err := this.createAnswer(answerCreate.Text, answerCreate.IsCorrect, i+1)
		if err != nil {
			return nil, err
		}

		answerEntities = append(answerEntities, answerEntity)
	}

	return &Question{
		id:        identificators.QuestionID(uuid.New()),
		byTopicID: topic,
		text:      textVO,
		answers:   answerEntities,
	}, nil
}

func (QuestionFabric) createAnswer(text string, isCorrect bool, serialNumber int) (*Answer, error) {
	answerText, err := NewAnswertext(text)
	if err != nil {
		return nil, err
	}

	if serialNumber < 1 {
		return nil, fmt.Errorf("Answer serial number cannot be less 1")
	}

	return &Answer{
		text:         answerText,
		isCorrect:    isCorrect,
		serialNumber: serialNumber,
	}, nil
}

var Fabric = QuestionFabric{}
