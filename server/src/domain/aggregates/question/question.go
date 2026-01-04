package question

import (
	"adaptivetesting/src/domain/aggregates/identificators"
)

type Question struct {
	id      identificators.QuestionID
	topicID identificators.TopicID
	text    QuestionText
	answers []Answer
}

func (this *Question) ID() identificators.QuestionID {
	return this.id
}

func (this *Question) TopicID() identificators.TopicID {
	return this.topicID
}

func (this *Question) Text() QuestionText {
	return this.text
}

func (this *Question) Answers() []Answer {
	return this.answers
}
