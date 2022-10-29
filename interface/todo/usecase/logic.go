package usecase

import (
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	"github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/errors"
)

type TodoUseCase struct {
	data domain.TodoData
}

func New(d domain.TodoData) *TodoUseCase {
	return &TodoUseCase{
		data: d,
	}
}

func (u *TodoUseCase) CreateNewTodo(req req.CreateNewTodo) (res *entities.Todo, err error, errType error) {
	//Validate request

	//Insert new todo data to db
	resData, err := u.data.CreateNew(req)
	if err != nil {
		return nil, err, errors.FailedCreateNewTodo
	}

	return resData, nil, nil
}
func (u *TodoUseCase) DeleteExistingTodo(req string) {}
func (u *TodoUseCase) GetAllListTodo(req string)     {}
func (u *TodoUseCase) GetDetailTodo(req string)      {}
func (u *TodoUseCase) UpdateExistingTodo(req string) {}
