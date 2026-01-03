package dto

type CourseResponse struct {
	ID          string
	Name        string
	CreatedByID string
}

type CourseResponseWithUser struct {
	ID        string
	Name      string
	CreatedBy UserResponse
}
