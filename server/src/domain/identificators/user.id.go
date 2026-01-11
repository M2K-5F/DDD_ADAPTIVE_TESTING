package identificators

import "github.com/google/uuid"

type UserID uuid.UUID

func (i UserID) String() string {
	return uuid.UUID(i).String()
}

func NewUserID() UserID {
	return UserID(uuid.New())
}

func ParseUserID(plain string) (UserID, error) {
	id, err := uuid.Parse(plain)
	if err != nil {
		return UserID(uuid.Nil), err
	}
	return UserID(id), nil
}
