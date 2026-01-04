package responses

type TopicResponseWithCourseAndUser struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	ByCourse  CourseResponse `json:"by_course"`
	CreatedBy UserResponse   `json:"created_by"`
}

type TopicResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ByCourseID  string `json:"by_course_id"`
	CreatedByID string `json:"created_by_id"`
}
