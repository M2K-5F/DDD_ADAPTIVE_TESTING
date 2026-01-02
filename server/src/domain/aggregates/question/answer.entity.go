package question

type Answer struct {
	serialNumber int
	text         AnswerText
	isCorrect    bool
}

func NewAnswer(serialNumber int, text string, isCorrect bool) (*Answer, error) {
	textVO, err := NewAnswertext(text)

	if err != nil {
		return nil, err
	}

	return &Answer{
		isCorrect:    isCorrect,
		text:         textVO,
		serialNumber: serialNumber,
	}, nil
}
