package request

type UpdateTaskRequest struct {
	Id          int    `json:"id"`
	Title       string `validate:"required,min=1,max=25" json:"title"`
	Description string `validate:"required,min=1,max=200" json:"description"`
	DueDate     string `validate:"required" json:"due_date"`
}
