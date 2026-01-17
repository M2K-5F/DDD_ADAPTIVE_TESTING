package group_usecases

import "adaptivetesting/src/application/interfaces"

type DependsModule struct {
	CourseReader interfaces.ICourseReader
	Writer       interfaces.IWriter
}

type GroupUseCase struct {
	deps *DependsModule
}

func New(module *DependsModule) *GroupUseCase {
	return &GroupUseCase{
		deps: module,
	}
}
