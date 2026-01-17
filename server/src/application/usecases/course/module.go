package course_usercases

import "adaptivetesting/src/application/interfaces"

type DependsModule struct {
	Writer    interfaces.IWriter
	CourseRdr interfaces.ICourseReader
	UserRdr   interfaces.IUserReader
}

type CourseUseCases struct {
	deps *DependsModule
}

func New(module *DependsModule) *CourseUseCases {
	return &CourseUseCases{
		deps: module,
	}
}
