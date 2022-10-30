package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/data"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/dispatch"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/usecase"
	"gorm.io/gorm"
)

type WithDependency struct {
	App *fiber.App
	DB  gorm.DB
}

func ServicesImpl(d WithDependency) {
	dataInterface := data.New(d.DB)
	useCaseInterface := usecase.New(dataInterface)
	handlerInterface := dispatch.New(useCaseInterface, dataInterface)
	dispatch.Routes(d.App, handlerInterface)
}
