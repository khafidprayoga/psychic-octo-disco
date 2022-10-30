package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/errors"
	"github.com/khafidprayoga/psychic-octo-disco/utils"
)

type TodoUseCase struct {
	data domain.TodoData
}

func New(d domain.TodoData) *TodoUseCase {
	return &TodoUseCase{
		data: d,
	}
}

func (u *TodoUseCase) CreateNewTodo(req request.CreateNewTodo) (res *entities.Todo, httpCode int, err error, errType error) {
	//Validate request
	if err := utils.ValidateStruct[request.CreateNewTodo](req); err != nil {
		return nil, fiber.StatusBadRequest, err, errors.InvalidRequestBody
	}

	//Insert new todo data to db
	resData, err := u.data.CreateNew(req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err, errors.FailedCreateNewTodo
	}

	return resData, fiber.StatusCreated, nil, nil
}
func (u *TodoUseCase) DeleteExistingTodo(req string) {}
func (u *TodoUseCase) GetAllListTodo(req string)     {}
func (u *TodoUseCase) GetDetailTodo(req string)      {}
func (u *TodoUseCase) UpdateExistingTodo(req string) {}
