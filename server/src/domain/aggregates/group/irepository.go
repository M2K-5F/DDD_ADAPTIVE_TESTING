package group

import "adaptivetesting/src/domain/aggregates/identificators"

type IRepository interface {
	Save(*Group) error
	GetByID(id identificators.GroupID) (*Group, error)
}
