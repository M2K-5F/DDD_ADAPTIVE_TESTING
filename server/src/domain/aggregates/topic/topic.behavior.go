package topic

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"
)

func (this *Topic) Activate() error {
	if this.isArchived {
		this.isArchived = false
		return nil
	}
	return fmt.Errorf("topic already active")
}

func (this *Topic) Archivate() error {
	if !this.isArchived {
		this.isArchived = true
		return nil
	}
	return fmt.Errorf("topic already archived")
}

func (this *Topic) IsCreatedBy(userID identificators.UserID) bool {
	return this.createdByID == userID
}
