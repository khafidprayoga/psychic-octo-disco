package req

type UpdateExistingTodo struct {
	id       string `validate:"-"`
	Title    string `json:"title"`
	IsActive *bool  `json:"is_active"`
}

func (t *UpdateExistingTodo) SetId(id string) {
	t.id = id
}

func (t *UpdateExistingTodo) GetId() string {
	return t.id
}
