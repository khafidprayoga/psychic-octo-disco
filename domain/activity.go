package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	"github.com/khafidprayoga/psychic-octo-disco/http/req"
)

type ActivityHandler interface {
	DeleteExistingActivity() fiber.Handler
	GetDetailActivity() fiber.Handler
	GetAllActivity() fiber.Handler
	PostNewActivity() fiber.Handler
}

type ActivityUseCase interface {
	DetailActvity(id string) (res *entities.Activity, httpCode int, errType error, srvError int)
	DeleteActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int)
	ListActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int)
	UpdateActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int)
}

type ActivityData interface {
	CreateNewActivity(req req.CreateNewTodo) (data *entities.Activity, err error)
	DeleteActivityData(id string) error
	DetailActivityData(id string) (*entities.Activity, error)
	ListActivityData(activityId string) ([]entities.Activity, error)
	UpdateActivityData(data req.UpdateExistingTodo) (*entities.Activity, error)
	ValidateTodo(id string) error
}
