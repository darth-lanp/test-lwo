package request

type CreateTaskRequest struct {
	Title       string `validate:"required,min=1,max=25" json:"title"`
	Description string `validate:"required,min=1,max=255" json:"description"`
	DueDate     string `validate:"required" json:"due_date"`
}
