package domain

import "github.com/gofiber/fiber/v2"

// TodoHandler for handle incoming http request with interface todo
type TodoHandler interface {
	DeleteData() fiber.Handler
	GetDetail() fiber.Handler
	GetList() fiber.Handler
	PostNewData() fiber.Handler
	UpdateData() fiber.Handler
}

// TodoUseCase for handle todo interface business logic
type TodoUseCase interface {
	CreateNewTodo(req string)
	DeleteExistingTodo(req string)
	GetAllListTodo(req string)
	GetDetailTodo(req string)
	UpdateExistingTodo(req string)
}

// TodoData to interact with mysql db
type TodoData interface {
	DeleteTodo(id string) error
	DetailTodo(id string) error
	ListAllTodo() error
	UpdateTodo(id string) error
	ValidateTodo(id string) error
}
