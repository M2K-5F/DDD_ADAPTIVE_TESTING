package usecases

import "adaptivetesting/src/application/interfaces"

type RegisterUserUC struct {
	userRepo interfaces.UserRepository
}

func (RegisterUserUC) Fabric(userRepoImplementation interfaces.UserRepository) *RegisterUserUC {
	return &RegisterUserUC{
		userRepo: userRepoImplementation,
	}
}

type Role string

const (
	Teacher Role = "Teacher"
	Student Role = "Student"
)

type RegisterUserDTO struct {
	UserName string
	Role     Role
}

func (r *RegisterUserUC) Execute() {

}
