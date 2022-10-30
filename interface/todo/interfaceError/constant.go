package interfaceError

import "errors"

var (
	InvalidRequestBody       = errors.New("invalid request body data")
	FailedCreateNewTodo      = errors.New("failed to create new todo")
	FailedDeleteExistingTodo = errors.New("failed to delete todo")
	FailedUpdateExistingTodo = errors.New("failed update existing todo")
	DataNotFound             = errors.New("sorry, todo data does not exist")
)
