package requests

type Role string

const (
	Teacher Role = "Teacher"
	Student Role = "Student"
)

type RegisterUserRequest struct {
	UserName string
	Role     Role
	Password string
}

type AuthUserRequest struct {
	UserName string
	Password string
}
