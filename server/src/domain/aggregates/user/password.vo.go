package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) Verify(plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(plain)) == nil
}

func NewPassword(plain string) (Password, error) {
	if len(plain) < 5 {
		return "", fmt.Errorf("password < 5 chars")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Password(hashed), nil
}
