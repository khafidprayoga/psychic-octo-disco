package req

type UpdateExistingTodo struct {
	id       string `validate:"-"`
	Title    string `json:"title" validate:"required"`
	IsActive *bool  `json:"is_active" validate:"required"`
}

func (t *UpdateExistingTodo) SetId(id string) {
	t.id = id
}

func (t *UpdateExistingTodo) GetId() string {
	return t.id
}
