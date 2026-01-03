package user

type Role string

const (
	TeacherRole Role = "TEACHER"
	StudentRole Role = "STUDENT"
)

func (r Role) String() string {
	return string(r)
}
