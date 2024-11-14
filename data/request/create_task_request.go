package request

import "time"

type CreateTaskRequest struct {
	Title       string    `validate:"required,min=1,max=200" json:"title"`
	Description string    `validate:"required,min=1,max=200" json:"description"`
	DueDate     time.Time `validate:"required,min=1,max=200" json:"due_date"`
}
