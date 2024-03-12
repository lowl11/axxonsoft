package dto

type Task struct {
	Method  string            `json:"method" validate:"required"`
	URL     string            `json:"url" validate:"required"`
	Headers map[string]string `json:"headers"`
}

type TaskResponse struct {
	StatusCode  int
	ContentType string
	Body        []byte
}
