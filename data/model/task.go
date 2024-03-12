package model

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID         `json:"id"`
	Status      string            `json:"status"`
	StatusCode  *int              `json:"httpStatusCode"`
	Headers     map[string]string `json:"headers"`
	Length      *int              `json:"length"`
	ErrorReason *string           `json:"error_reason,omitempty"`
}
