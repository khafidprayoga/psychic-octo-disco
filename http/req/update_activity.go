package req

type UpdateExistingActivity struct {
	id    string `validate:"-"`
	Title string `json:"title" validate:"required"`
}

func (t *UpdateExistingActivity) SetId(id string) {
	t.id = id
}

func (t *UpdateExistingActivity) GetId() string {
	return t.id
}
