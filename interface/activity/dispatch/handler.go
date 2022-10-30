package dispatch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
)

type ActivityHandler struct {
	useCaseImpl domain.ActivityUseCase
	dataImpl    domain.ActivityData
}

func New(uc domain.ActivityUseCase, data domain.ActivityData) domain.ActivityHandler {
	return &ActivityHandler{
		useCaseImpl: uc,
		dataImpl:    data,
	}
}

func (h *ActivityHandler) DeleteExistingActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func (h *ActivityHandler) GetDetailActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func (h *ActivityHandler) GetAllActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func (h *ActivityHandler) PostNewActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
