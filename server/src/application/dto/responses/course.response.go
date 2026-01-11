package responses

type CourseResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatedByID string `json:"created_by_id"`
}

type CourseNestedResponse struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	CreatedBy UserResponse `json:"created_by"`
}
