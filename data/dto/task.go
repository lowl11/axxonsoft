package dto

type Task struct {
	Method  string            `json:"method" validate:"required"`
	URL     string            `json:"url" validate:"required"`
	Headers map[string]string `json:"headers"`
	Body    map[string]any    `json:"body"`
}
