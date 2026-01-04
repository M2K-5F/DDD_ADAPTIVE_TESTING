package identificators

import "github.com/google/uuid"

type GroupID uuid.UUID

func (id GroupID) String() string {
	return uuid.UUID(id).String()
}
