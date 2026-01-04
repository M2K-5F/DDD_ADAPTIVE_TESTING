package responses

type QuestionResponse struct {
	ID        string           `json:"id"`
	Text      string           `json:"text"`
	ByTopicID string           `json:"by_topic_id"`
	Answers   []AnswerResponse `json:"answers"`
}

type AnswerResponse struct {
	SerialNumber int    `json:"serial_number"`
	Text         string `json:"text"`
	IsCorrect    bool   `json:"is_correct"`
}

type QuestionToPassResponse struct {
	ID        string                 `json:"id"`
	Text      string                 `json:"text"`
	ByTopicID string                 `json:"by_topic_id"`
	Answers   []AnswerToPassResponse `json:"answers"`
}

type AnswerToPassResponse struct {
	SerialNumber int    `json:"serial_number"`
	Text         string `json:"text"`
}
