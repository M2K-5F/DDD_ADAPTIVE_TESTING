package question

type Answer struct {
	serialNumber int
	text         AnswerText
	isCorrect    bool
}

func (this *Answer) SerialNumber() int {
	return this.serialNumber
}

func (this *Answer) Text() AnswerText {
	return this.text
}

func (this *Answer) IsCorrect() bool {
	return this.isCorrect
}
