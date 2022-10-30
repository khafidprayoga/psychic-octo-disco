package activity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity/data"
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity/dispatch"
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity/usecase"
	"gorm.io/gorm"
)

type WithDependency struct {
	App *fiber.App
	DB  gorm.DB
}

func ServiceImpl(d WithDependency) {
	dataInterface := data.New(d.DB)
	useCaseInterface := usecase.New(dataInterface)
	handlerInterface := dispatch.New(useCaseInterface, dataInterface)
	dispatch.Routes(d.App, handlerInterface)
}
