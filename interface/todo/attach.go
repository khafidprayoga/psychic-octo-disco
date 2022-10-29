package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/data"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/dispatch"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/usecase"
	"gorm.io/gorm"
)

type InitServicesTodo struct {
	App *fiber.App
	Db  gorm.DB
}

func ServicesImpl(d InitServicesTodo) {
	dataInterface := data.New(d.Db)
	useCaseInterface := usecase.New(dataInterface)
	handlerInterface := dispatch.New(useCaseInterface, dataInterface)
	dispatch.Routes(d.App, handlerInterface)
}
