package responses

type TopicResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ByCourseID  string `json:"by_course_id"`
	CreatedByID string `json:"created_by_id"`
}

type TopicNestedResponse struct {
	ID        string
	Name      string
	ByCourse  CourseResponse
	CreatedBy UserResponse
}
