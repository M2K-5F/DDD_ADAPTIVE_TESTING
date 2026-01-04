package topic_usecases

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/dto/responses"
	"adaptivetesting/src/application/mappers"
	"adaptivetesting/src/domain/aggregates/course"
	"adaptivetesting/src/domain/aggregates/topic"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
)

type CreateTopic struct {
	topicRepo  topic.ITopicRepository
	courseRepo course.ICourseRepository
}

func (this CreateTopic) Execute(current_user *user.User, data requests.CreateTopicRequest) (*responses.TopicResponseWithCourseAndUser, error) {
	if !current_user.IsTeacher() {
		return nil, fmt.Errorf("Only teachers can create topics")
	}

	current_course, err := this.courseRepo.GetById(data.CourseId)
	if err != nil {
		return nil, err
	}

	if !current_course.IsUserCreator(current_user) {
		return nil, fmt.Errorf("Only creator of this course can add topic to it")
	}

	createdTopic, err := topic.NewTopic(data.Name, *current_course, *current_user)
	if err != nil {
		return nil, err
	}

	if err := this.topicRepo.Save(createdTopic); err != nil {
		return nil, err
	}

	if err := current_course.AddTopic(createdTopic.ID()); err != nil {
		return nil, err
	}

	if err := this.courseRepo.Save(current_course); err != nil {
		return nil, err
	}

	return mappers.TopicToNestedResponse(createdTopic, current_course, current_user), nil
}
