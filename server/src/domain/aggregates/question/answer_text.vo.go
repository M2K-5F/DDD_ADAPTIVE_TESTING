package question

import "fmt"

type AnswerText string

func (a AnswerText) String() string {
	return string(a)
}

func NewAnswertext(plain string) (AnswerText, error) {
	if len(plain) < 16 {
		return "", fmt.Errorf(
			"Text length (%d) is less than required minimum (16)",
			len(plain),
		)
	}

	if len(plain) > 256 {
		return "", fmt.Errorf(
			"Text length (%d) is greater than required maximum (256)",
			len(plain),
		)
	}
	return AnswerText(plain), nil
}
