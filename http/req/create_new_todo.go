package req

type CreateNewTodo struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupID *int   `json:"activity_group_id" validate:"required,numeric,min=1"`
}
