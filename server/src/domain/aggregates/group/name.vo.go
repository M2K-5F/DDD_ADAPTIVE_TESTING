package group

import "fmt"

type GroupName string

func (g GroupName) String() string {
	return string(g)
}

func NewGroupName(plain string) (GroupName, error) {
	if len(plain) < 8 {
		return "", fmt.Errorf(
			"groupName length (%d) is less than required minimum (8)",
			len(plain),
		)
	}

	if len(plain) > 64 {
		return "", fmt.Errorf(
			"GroupName length (%d) is greater than required maximum (64)",
			len(plain),
		)
	}
	return GroupName(plain), nil
}
