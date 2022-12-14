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
	PatchExistingActivity() fiber.Handler
}

type ActivityUseCase interface {
	CreateNewActivity(req req.CreateNewActivity) (res *entities.Activity, httpCode int, errType error, srvError int)
	DetailActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int)
	DeleteActivity(id string) (httpCode int, errType error, srvError int)
	ListActivity() (res []entities.Activity, httpCode int, errType error, srvError int)
	UpdateActivity(req req.UpdateExistingActivity) (res *entities.Activity, httpCode int, errType error, srvError int)
}

type ActivityData interface {
	CreateNewActivity(req req.CreateNewActivity) (data *entities.Activity, err error)
	DeleteActivityData(id string) error
	DetailActivityData(id string) (*entities.Activity, error)
	ListActivityData() ([]entities.Activity, error)
	UpdateActivityData(req req.UpdateExistingActivity) (data *entities.Activity, err error)
	ValidateActivity(id string) error
}
