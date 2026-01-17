package topic_usecases

import "adaptivetesting/src/application/interfaces"

type DependsModule struct {
	Writer    interfaces.IWriter
	CourseRdr interfaces.ICourseReader
}

type TopicUseCases struct {
	deps *DependsModule
}

func New(module *DependsModule) *TopicUseCases {
	return &TopicUseCases{
		deps: module,
	}
}
