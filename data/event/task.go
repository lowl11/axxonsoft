package event

import "github.com/google/uuid"

type CallTask struct {
	TaskID uuid.UUID `json:"task_id" validate:"required,uuid"`
}
