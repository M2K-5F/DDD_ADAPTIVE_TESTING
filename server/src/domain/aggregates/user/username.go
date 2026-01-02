package user

import "fmt"

type Username string

func (u Username) String() string {
	return string(u)
}

func NewUsername(username string) (Username, error) {
	if len(username) < 5 {
		return "", fmt.Errorf(
			"username length (%d) is less than required minimum (5)",
			len(username),
		)
	}

	if len(username) > 32 {
		return "", fmt.Errorf(
			"username length (%d) is greater than required maximum (32)",
			len(username),
		)
	}
	return Username(username), nil
}
