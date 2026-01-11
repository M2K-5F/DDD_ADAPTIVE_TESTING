package question

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"
)

type Question struct {
	id         identificators.QuestionID
	byTopicID  identificators.TopicID
	byCourseID identificators.CourseID
	text       QuestionText
	answers    []*Answer
}

func (this *Question) ID() identificators.QuestionID {
	return this.id
}

func (this *Question) ByCourseID() identificators.CourseID {
	return this.byCourseID
}

func (this *Question) TopicID() identificators.TopicID {
	return this.byTopicID
}

func (this *Question) Text() QuestionText {
	return this.text
}

func (this *Question) Answers() []*Answer {
	return this.answers
}

func (this *Question) GetAnswerBySerial(number int) (*Answer, error) {
	for _, answer := range this.answers {
		if answer.serialNumber == number {
			return answer, nil
		}
	}
	return nil, fmt.Errorf("Answer with serial (%d) not found", number)
}
