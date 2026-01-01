package user

import (
	valueobjects "adaptivetesting/src/domain/user/value_objects"
	"time"

	"github.com/google/uuid"
)

type UserID uuid.UUID

type User struct {
	ID           UserID
	RegisteredAt time.Time
	UserName     string
	Roles        []valueobjects.Role
}
