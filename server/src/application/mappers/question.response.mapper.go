package mappers

import (
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/domain/aggregates/question"
)

func QuestionToResponse(question *question.Question) *responses.QuestionResponse {
	var answers []responses.AnswerResponse

	for _, answer := range question.Answers() {
		answers = append(answers, responses.AnswerResponse{
			Text:         answer.Text().String(),
			SerialNumber: answer.SerialNumber(),
			IsCorrect:    answer.IsCorrect(),
		})
	}

	return &responses.QuestionResponse{
		ID:        question.ID().String(),
		Text:      question.Text().String(),
		ByTopicID: question.TopicID().String(),
		Answers:   answers,
	}
}

func QuestionToResponseToPass(question *question.Question) *responses.QuestionToPassResponse {
	var answers []responses.AnswerToPassResponse

	for _, answer := range question.Answers() {
		answers = append(answers, responses.AnswerToPassResponse{
			Text:         answer.Text().String(),
			SerialNumber: answer.SerialNumber(),
		})
	}

	return &responses.QuestionToPassResponse{
		ID:        question.ID().String(),
		Text:      question.Text().String(),
		ByTopicID: question.TopicID().String(),
		Answers:   answers,
	}
}
