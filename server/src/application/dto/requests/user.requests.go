package requests

type Role string

const (
	Teacher Role = "Teacher"
	Student Role = "Student"
)

type RegisterUserDTO struct {
	UserName string
	Role     Role
	Password string
}

type AuthUserDTO struct {
	UserName string
	Password string
}
