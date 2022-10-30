package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/interfaceError"
	"github.com/khafidprayoga/psychic-octo-disco/utils"
	"log"
)

type TodoUseCase struct {
	data domain.TodoData
}

func New(d domain.TodoData) *TodoUseCase {
	return &TodoUseCase{
		data: d,
	}
}

func (u *TodoUseCase) CreateNewTodo(req request.CreateNewTodo) (res *entities.Todo, httpCode int, errType error, srvError int) {
	//Validate request
	if err := utils.ValidateStruct[request.CreateNewTodo](req); err != nil {
		return nil, fiber.StatusBadRequest, interfaceError.InvalidRequestBody, utils.HTTPRequestErr
	}

	//Insert new todo data to db
	resData, err := u.data.CreateNew(req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err, utils.DatabaseError
	}

	return resData, fiber.StatusCreated, nil, 0
}

func (u *TodoUseCase) DeleteExistingTodo(id string) (httpCode int, errType error, srvError int) {
	// Validate tood exist
	if err := u.data.ValidateTodo(id); err != nil {
		return fiber.StatusNotFound, err, utils.HTTPRequestErr
	}

	if err := u.data.DeleteTodo(id); err != nil {
		log.Println(err)
		return fiber.StatusInternalServerError, err, utils.DatabaseError

	}

	return fiber.StatusOK, nil, 0
}

func (u *TodoUseCase) GetAllListTodo(req string)     {}
func (u *TodoUseCase) GetDetailTodo(req string)      {}
func (u *TodoUseCase) UpdateExistingTodo(req string) {}
