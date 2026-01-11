package question_usercases

import (
	"adaptivetesting/src/application/dto/mappers"
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/interfaces"
	"adaptivetesting/src/domain/aggregates/question"
	"adaptivetesting/src/domain/aggregates/user"
	"context"
	"fmt"
)

type CreateQuestion struct {
	questionReader interfaces.IQuestionReader
	topicReader    interfaces.ITopicReader
	writer         interfaces.IWriter
}

func (this CreateQuestion) Execute(
	ctx context.Context,
	current_user *user.User,
	data *requests.CreateQuestionRequest,
) (
	*responses.QuestionResponse,
	error,
) {
	if !current_user.IsTeacher() {
		return nil, fmt.Errorf("Only teachers can create a questions")
	}

	current_topic, err := this.topicReader.GetByID(ctx, data.TopicID)
	if err != nil {
		return nil, err
	}

	if !current_topic.IsCreatedBy(current_user.ID()) {
		return nil, fmt.Errorf("Only topic creators can add questions to this topic")
	}

	var answers []*question.AnswerCreate
	for _, ans := range data.Answers {
		if err != nil {
			return nil, err
		}
		answers = append(
			answers,
			&question.AnswerCreate{
				Text:      ans.Text,
				IsCorrect: ans.IsCorrect,
			},
		)
	}

	createdQuestion, err := question.Fabric.CreateQuestion(current_topic.ID(), data.Text, answers)
	if err != nil {
		return nil, err
	}

	if err := this.writer.Execute(ctx, func(writer interfaces.ITransactionWriter) error {
		return writer.SaveQuestion(createdQuestion)
	}); err != nil {
		return nil, err
	}

	return mappers.QuestionToResponse(createdQuestion), nil
}
