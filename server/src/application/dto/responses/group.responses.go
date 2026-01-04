package responses

type GroupResponse struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	CreatedByID  string `json:"created_by_id"`
	ByCourseID   string `json:"by_course_id"`
	StudentCount int    `json:"student_count"`
}

type GroupNestedResponse struct {
	Name         string         `json:"name"`
	ID           string         `json:"id"`
	CreatedBy    UserResponse   `json:"created_by"`
	ByCourse     CourseResponse `json:"by_course"`
	StudentCount int            `json:"student_count"`
}
