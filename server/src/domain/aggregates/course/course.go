package course

import (
	"adaptivetesting/src/domain/aggregates/identificators"
	"adaptivetesting/src/domain/aggregates/user"
	"fmt"
	"slices"
)

type Course struct {
	id          identificators.CourseID
	createdByID identificators.UserID
	name        CourseName
	isArchived  bool

	topicCount int
	topicIDs   []identificators.TopicID
	groupCount int
	groupIDs   []identificators.GroupID
}

// QUERY
func (c *Course) ID() identificators.CourseID {
	return c.id
}

func (c *Course) Name() CourseName {
	return c.name
}

func (c *Course) CreatedByID() identificators.UserID {
	return c.createdByID
}

func (c *Course) IsArchived() bool {
	return c.isArchived
}

func (c *Course) TopicCount() int {
	return c.topicCount
}

func (c *Course) GroupCount() int {
	return c.groupCount
}

func (this *Course) IsTopicPresentAtCourse(topicID identificators.TopicID) bool {
	return slices.Contains(this.topicIDs, topicID)
}

func (this *Course) IsGroupPresentAtCourse(groupID identificators.GroupID) bool {
	return slices.Contains(this.groupIDs, groupID)
}

// COMMAND
func (this *Course) AddTopic(topicID identificators.TopicID) error {
	if this.IsTopicPresentAtCourse(topicID) {
		return fmt.Errorf("This topic already presented at course")
	}

	this.topicCount += 1
	this.topicIDs = append(this.topicIDs, topicID)
	return nil
}

func (this *Course) AddGroup(groupID identificators.GroupID) error {
	if this.IsGroupPresentAtCourse(groupID) {
		return fmt.Errorf("Group Already presented at course")
	}
	this.groupCount += 1
	this.groupIDs = append(this.groupIDs, groupID)
	return nil
}

func (this *Course) RemoveGroup() error {
	if this.groupCount > 1 {
		this.groupCount -= 1
		return nil
	}
	return fmt.Errorf("course group count is 0")
}

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

func (this *Course) IsUserCreator(user *user.User) bool {
	return this.createdByID == user.ID()
}
