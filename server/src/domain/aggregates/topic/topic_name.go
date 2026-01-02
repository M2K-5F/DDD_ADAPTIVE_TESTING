package topic

import "fmt"

type TopicName string

func (t TopicName) String() string {
	return string(t)
}

func NewTopicName(plain string) (TopicName, error) {
	if len(plain) < 5 {
		return "", fmt.Errorf(
			"name length (%d) is less than required minimum (5)",
			len(plain),
		)
	}

	if len(plain) > 32 {
		return "", fmt.Errorf(
			"Name length (%d) is greater than required maximum (32)",
			len(plain),
		)
	}
	return TopicName(plain), nil
}
