package errors

import "errors"

var (
	InvalidRequestBody  = errors.New("invalid request body data")
	FailedCreateNewTodo = errors.New("failed to create new todo")
	DataNotFound        = errors.New("sorry, todo data not exist")
)
