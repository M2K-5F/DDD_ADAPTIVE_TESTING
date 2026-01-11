package course

import "fmt"

type CourseName string

func (c CourseName) String() string {
	return string(c)
}

func NewCourseName(plain string) (CourseName, error) {
	if len(plain) < 5 {
		return "", fmt.Errorf(
			"CourseName length (%d) is less than required minimum (5)",
			len(plain),
		)
	}
	if len(plain) > 64 {
		return "", fmt.Errorf(
			"CourseName length (%d) is greater than required maximum (64)",
			len(plain),
		)
	}
	return CourseName(plain), nil
}
