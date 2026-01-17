package question_usercases

import "adaptivetesting/src/application/interfaces"

type DependsModule struct {
	QuestionReader interfaces.IQuestionReader
	TopicReader    interfaces.ITopicReader
	Writer         interfaces.IWriter
}

type QuestionUseCases struct {
	deps *DependsModule
}

func New(module *DependsModule) *QuestionUseCases {
	return &QuestionUseCases{
		deps: module,
	}
}
