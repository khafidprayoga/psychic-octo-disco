package usecase

import (
	"fmt"
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
		errMsg := fmt.Errorf("%v, reason: %v", interfaceError.InvalidRequestBody, err)
		return nil, fiber.StatusBadRequest, errMsg, utils.HTTPRequestErr
	}

	//Insert new todo data to db
	resData, err := u.data.CreateNew(req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err, utils.DatabaseError
	}

	return resData, fiber.StatusCreated, nil, 0
}

func (u *TodoUseCase) DeleteExistingTodo(id string) (httpCode int, errType error, srvError int) {
	// Validate todo exist
	if err := u.data.ValidateTodo(id); err != nil {
		return fiber.StatusNotFound, err, utils.HTTPRequestErr
	}

	if err := u.data.DeleteTodo(id); err != nil {
		log.Println(err)
		return fiber.StatusInternalServerError, err, utils.DatabaseError

	}

	return fiber.StatusOK, nil, 0
}

func (u *TodoUseCase) GetAllListTodo(activityId string) (res []entities.Todo, httpCode int, errType error, srvError int) {
	listTodo, err := u.data.ListAllTodo(activityId)
	if err != nil {
		return []entities.Todo{}, fiber.StatusInternalServerError, interfaceError.FailedGetListTodo, utils.DatabaseError
	}

	return listTodo, fiber.StatusOK, nil, 0
}

func (u *TodoUseCase) GetDetailTodo(id string) (res *entities.Todo, httpCode int, errType error, srvError int) {
	// Validate todo exist
	if err := u.data.ValidateTodo(id); err != nil {
		return nil, fiber.StatusNotFound, interfaceError.DataNotFound, utils.HTTPRequestErr
	}
	entitiesData, err := u.data.DetailTodo(id)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err, utils.DatabaseError
	}
	return entitiesData, fiber.StatusOK, nil, 0
}

func (u *TodoUseCase) UpdateExistingTodo(req request.UpdateExistingTodo) (res *entities.Todo, httpCode int, errType error, srvError int) {
	//Validate request
	if err := utils.ValidateStruct[request.UpdateExistingTodo](req); err != nil {
		errMsg := fmt.Errorf("%v, reason: %v", interfaceError.InvalidRequestBody, err)
		return nil, fiber.StatusBadRequest, errMsg, utils.HTTPRequestErr
	}

	// Validate todo exist
	if err := u.data.ValidateTodo(req.GetId()); err != nil {
		return nil, fiber.StatusNotFound, err, utils.NotFoundErr
	}

	updatedEntities, err := u.data.UpdateTodo(req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, interfaceError.FailedUpdateExistingTodo, utils.DatabaseError
	}

	return updatedEntities, fiber.StatusOK, nil, 0
}
