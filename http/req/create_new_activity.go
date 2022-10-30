package req

type CreateNewActivity struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
