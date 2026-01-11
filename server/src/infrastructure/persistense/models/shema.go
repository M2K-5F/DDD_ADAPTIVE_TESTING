package models

import (
	"time"
)

type UserRows struct {
	ID           string    `db:"id"`
	RegisteredAt time.Time `db:"registered_at"`
	UserName     string    `db:"user_name"`
	PasswordHash string    `db:"password_hash"`
	Roles        []string  `db:"roles"`
}

type CourseRows struct {
	ID          string
	CreatedByID string
	Name        string
	IsArchived  bool
}

type TopicRows struct {
	ID            string
	CreatedByID   string
	ByCourseID    string
	Name          string
	IsArchived    bool
	QuestionCount int
}

type QuestionRows struct {
	ID        string
	ByTopicID string
	Text      string
}

type AnswerRows struct {
	ID           string
	ByQuestionID string
	Text         string
	IsCorrect    bool
	SerialNumber int
}

type GroupRows struct {
	ID              string
	ByCourseID      string
	CreatedByID     string
	Name            string
	StudentCount    int
	MaxStudentCount int
}

type EnrollmentRows struct {
	ID       string
	GroupID  string
	CourseID string
	UserID   string
	Progress float32
}

type TopicProgressRows struct {
	ID           string
	EnrollmentID string
	MaxScore     float32
}

type TopicAttemptRows struct {
	ID              string
	TopicProgressID string
	Score           float32
	AttemptedAt     time.Time
}
