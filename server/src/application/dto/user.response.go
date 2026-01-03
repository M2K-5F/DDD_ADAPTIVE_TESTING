package dto

import "time"

type UserResponse struct {
	ID           string
	UserName     string
	Roles        []string
	RegisteredAt time.Time
}
