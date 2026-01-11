package course

import (
	"adaptivetesting/src/domain/identificators"
	"fmt"
)

func (this *Course) Activate() error {
	if this.isArchived {
		this.isArchived = false
		return nil
	}
	return fmt.Errorf("course already active")
}

func (this *Course) Archivate() error {
	if !this.isArchived {
		this.isArchived = true
		return nil
	}
	return fmt.Errorf("course already archived")
}

func (this *Course) IsCreatedBy(userID identificators.UserID) bool {
	return this.createdByID == userID
}
