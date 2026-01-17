package question

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"

	"github.com/google/uuid"
)

type AnswerCreate struct {
	IsCorrect    bool
	Text         string
	SerialNumber int
}

type QuestionFabric struct{}

func (this QuestionFabric) CreateQuestion(topicID identificators.TopicID, courseID identificators.CourseID, text string, answers []*AnswerCreate) (*Question, error) {
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
		id:         identificators.QuestionID(uuid.New()),
		byTopicID:  topicID,
		byCourseID: courseID,
		text:       textVO,
		answers:    answerEntities,
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

func (this *QuestionFabric) Recover(id, courseID, topicID, text string, answersCreate []*AnswerCreate) (*Question, error) {
	var answers []*Answer

	for _, answerCreate := range answersCreate {
		answer, err := this.createAnswer(answerCreate.Text, answerCreate.IsCorrect, answerCreate.SerialNumber)
		if err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	topicIDVO, err := identificators.ParseTopicID(topicID)
	if err != nil {
		return nil, err
	}

	courseIDVO, err := identificators.ParseCourseID(courseID)
	if err != nil {
		return nil, err
	}

	idVO, err := identificators.ParseQuestionID(id)
	if err != nil {
		return nil, err
	}

	textVO, err := NewQuestionText(text)
	if err != nil {
		return nil, err
	}

	question := Question{
		id:         idVO,
		byCourseID: courseIDVO,
		byTopicID:  topicIDVO,
		text:       textVO,
		answers:    answers,
	}

	return &question, nil
}

var Fabric = QuestionFabric{}
