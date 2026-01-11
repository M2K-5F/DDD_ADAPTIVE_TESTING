package identificators

import "github.com/google/uuid"

type GroupID uuid.UUID

func (id GroupID) String() string {
	return uuid.UUID(id).String()
}

func NewGroupID() GroupID {
	return GroupID(uuid.New())
}

func ParseGroupID(plain string) (GroupID, error) {
	id, err := uuid.Parse(plain)
	if err != nil {
		return GroupID(uuid.Nil), err
	}
	return GroupID(id), nil
}
