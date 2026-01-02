package interfaces

import "adaptivetesting/src/domain/aggregates"

type UserRepository interface {
	Create(user aggregates.User) error
	GetByID(id aggregates.UserID) (aggregates.User, error)
	Update(user aggregates.User) error
}
