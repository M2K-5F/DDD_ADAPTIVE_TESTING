package question

import "fmt"

type QuestionText string

func (q QuestionText) String() string {
	return string(q)
}

func NewQuestionText(plain string) (QuestionText, error) {
	if len(plain) < 16 {
		return "", fmt.Errorf(
			"Text length (%d) is less than required minimum (16)",
			len(plain),
		)
	}

	if len(plain) > 256 {
		return "", fmt.Errorf(
			"Taxt length (%d) is greater than required maximum (256)",
			len(plain),
		)
	}
	return QuestionText(plain), nil
}
