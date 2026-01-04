package question_usercases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type CreateQuestion struct {
	questionRepo question.IRepository
	topicRepo    topic.ITopicRepository
}

func (this CreateQuestion) Execute(current_user *user.User, data *requests.CreateQuestionRequest) (*responses.QuestionResponse, error) {
	if !current_user.IsTeacher() {
		return nil, fmt.Errorf("Only teachers can create a questions")
	}

	current_topic, err := this.topicRepo.GetByID(data.TopicID)
	if err != nil {
		return nil, err
	}

	if !current_topic.IsCreatedBy(current_user) {
		return nil, fmt.Errorf("Only topic creators can add questions to this topic")
	}

	var answers []question.AnswerDTO
	for _, answer := range data.Answers {
		answers = append(answers, question.AnswerDTO{
			Text:      answer.Text,
			IsCorrect: answer.IsCorrect,
		})
	}

	createdQuestion, err := question.NewQuestion(*current_topic, data.Text, answers)
	if err != nil {
		return nil, err
	}

	if err := current_topic.AddQuestion(); err != nil {
		return nil, err
	}

	if err := this.questionRepo.Save(createdQuestion); err != nil {
		return nil, err
	}

	if err := this.topicRepo.Save(current_topic); err != nil {
		return nil, err
	}

	return mappers.QuestionToResponse(createdQuestion), nil
}

func Fabric(topicRepo topic.ITopicRepository, questionRepo question.IRepository) (*CreateQuestion, error) {
	return &CreateQuestion{
		topicRepo:    topicRepo,
		questionRepo: questionRepo,
	}, nil
}
