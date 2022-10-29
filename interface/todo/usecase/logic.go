package usecase

import "github.com/khafidprayoga/psychic-octo-disco/domain"

type TodoUseCase struct {
	data domain.TodoData
}

func New(d domain.TodoData) *TodoUseCase {
	return &TodoUseCase{
		data: d,
	}
}

func (u *TodoUseCase) CreateNewTodo(req string)      {}
func (u *TodoUseCase) DeleteExistingTodo(req string) {}
func (u *TodoUseCase) GetAllListTodo(req string)     {}
func (u *TodoUseCase) GetDetailTodo(req string)      {}
func (u *TodoUseCase) UpdateExistingTodo(req string) {}
