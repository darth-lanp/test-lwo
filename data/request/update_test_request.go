package request

import "time"

type UpdateTaskRequest struct {
	Id          int       `validate:"required"`
	Title       string    `validate:"required max=200,min=1" json:"title"`
	Description string    `validate:"required max=200,min=1" json:"description"`
	DueDate     time.Time `validate:"required max=200,min=1" json:"due_date"`
}
