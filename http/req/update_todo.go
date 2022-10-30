package req

type UpdateExistingTodo struct {
	id       int    `validate:"-"`
	Title    string `json:"title" validate:"required"`
	IsActive *bool  `json:"is_active" validate:"required"`
}

func (t *UpdateExistingTodo) SetId(id int) {
	t.id = id
}

func (t *UpdateExistingTodo) GetId() int {
	return t.id
}
