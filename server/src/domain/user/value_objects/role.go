package valueobjects

type RoleName string

const (
	TeacherRole RoleName = "TEACHER"
	StudentRole RoleName = "STUDENT"
)

type Role struct {
	Name RoleName
}
