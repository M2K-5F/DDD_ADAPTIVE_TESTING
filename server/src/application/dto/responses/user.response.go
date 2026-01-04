package responses

import "time"

type UserResponse struct {
	ID           string    `json:"id"`
	UserName     string    `json:"user_name"`
	Roles        []string  `json:"roles"`
	RegisteredAt time.Time `json:"registered_at"`
}
