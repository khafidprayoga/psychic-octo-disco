package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	"github.com/khafidprayoga/psychic-octo-disco/http/req"
)

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
	CreateNewTodo(req req.CreateNewTodo) (res *entities.Todo, httpCode int, errType error, srvError int)
	DeleteExistingTodo(id string) (httpCode int, errType error, srvError int)
	GetAllListTodo(req string)
	GetDetailTodo(id string) (res *entities.Todo, httpCode int, errType error, srvError int)
	UpdateExistingTodo(req req.UpdateExistingTodo) (res *entities.Todo, httpCode int, errType error, srvError int)
}

// TodoData to interact with mysql db
type TodoData interface {
	CreateNew(req req.CreateNewTodo) (data *entities.Todo, err error)
	DeleteTodo(id string) error
	DetailTodo(id string) (*entities.Todo, error)
	ListAllTodo() error
	UpdateTodo(data req.UpdateExistingTodo) (*entities.Todo, error)
	ValidateTodo(id string) error
}
