package user

import (
	"slices"
)

func (this *User) IsStudent() bool {
	if slices.Contains(this.roles, StudentRole) {
		return true
	}

	return false
}

func (this *User) IsTeacher() bool {
	if slices.Contains(this.roles, TeacherRole) {
		return true
	}

	return false
}

func (this *User) VerifyPassword(plain string) bool {
	return this.passwordHash.Verify(plain)
}
