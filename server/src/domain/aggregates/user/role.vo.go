package user

import "fmt"

type Role string

const (
	TeacherRole Role = "TEACHER"
	StudentRole Role = "STUDENT"
)

func (r Role) String() string {
	return string(r)
}

func NewRole(plain string) (Role, error) {
	if plain != "TEACHER" && plain != "STUDENT" {
		return "", fmt.Errorf("Role is not valid (%s)", plain)
	}

	if plain == "TEACHER" {
		return TeacherRole, nil
	} else {
		return StudentRole, nil
	}
}
