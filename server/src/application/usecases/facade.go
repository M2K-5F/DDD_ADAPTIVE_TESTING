package usecases

import (
	course_usercases "adaptivetesting/src/application/usecases/course"
	group_usecases "adaptivetesting/src/application/usecases/group"
	question_usercases "adaptivetesting/src/application/usecases/question"
	topic_usecases "adaptivetesting/src/application/usecases/topic"
	user_usercases "adaptivetesting/src/application/usecases/user"
)

type User = user_usercases.UserUseCases

type Course = course_usercases.CourseUseCases

type Topic = topic_usecases.TopicUseCases

type Question = question_usercases.QuestionUseCases

type Group = group_usecases.GroupUseCase
